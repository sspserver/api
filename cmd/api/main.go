package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/99designs/basicauth-go"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/goconfig"
	"github.com/demdxx/sendmsg"
	"github.com/demdxx/sendmsg/sender/email"
	"github.com/demdxx/sendmsg/sender/wrapper"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/geniusrabbit/blaze-api/pkg/auth"
	"github.com/geniusrabbit/blaze-api/pkg/auth/devtoken"
	"github.com/geniusrabbit/blaze-api/pkg/auth/elogin/facebook"
	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	"github.com/geniusrabbit/blaze-api/pkg/auth/oauth2"
	"github.com/geniusrabbit/blaze-api/pkg/context/ctxlogger"
	"github.com/geniusrabbit/blaze-api/pkg/context/version"
	"github.com/geniusrabbit/blaze-api/pkg/database"
	_ "github.com/geniusrabbit/blaze-api/pkg/gopentracing"
	"github.com/geniusrabbit/blaze-api/pkg/messanger"
	"github.com/geniusrabbit/blaze-api/pkg/migratedb"
	"github.com/geniusrabbit/blaze-api/pkg/permissions"
	"github.com/geniusrabbit/blaze-api/pkg/profiler"
	"github.com/geniusrabbit/blaze-api/pkg/zlogger"
	"github.com/geniusrabbit/blaze-api/repository/historylog/middleware/gormlog"
	"github.com/geniusrabbit/blaze-api/repository/socialauth/delivery/rest"

	statsClient "github.com/geniusrabbit/archivarius/client"
	billingClient "github.com/geniusrabbit/billing/client"

	"github.com/sspserver/api/cmd/api/appcontext"
	"github.com/sspserver/api/cmd/api/appinit"
	"github.com/sspserver/api/cmd/api/server"
	rtbsourceuc "github.com/sspserver/api/internal/repository/rtbsource/usecase"
	"github.com/sspserver/api/internal/repository/statistic"
	"github.com/sspserver/api/internal/server/graphql"
	"github.com/sspserver/api/internal/server/graphql/resolvers"
	"github.com/sspserver/api/private/emails"
)

var (
	buildDate    = ""
	buildCommit  = ""
	buildVersion = "develop"
)

func init() {
	runMigrations := flag.Bool("run-migrations", false, "Run database migrations")
	flag.Parse()

	conf := &appcontext.Config
	fatalError(goconfig.Load(conf), "load config:")

	if conf.IsDebug() || conf.IsInfo() {
		fmt.Println(conf)
	}

	// Migrate database schemas
	if *runMigrations {
		fmt.Println("Run database migrations")
		fatalError(migratedb.Migrate(conf.System.Storage.MasterConnect, []migratedb.MigrateSource{
			{
				URI:                   []string{"file:///data/migrations/initial"},
				SchemaMigrationsTable: "schema_migrations_initial",
			},
			// {
			// 	URI:                   []string{"file:///data/migrations/fixtures"},
			// 	SchemaMigrationsTable: "schema_migrations_dev",
			// },
			{
				URI:                   []string{"file:///data/migrations/project"},
				SchemaMigrationsTable: "schema_migrations",
			},
		}), "migrate database")
	}
}

func initZapLogger() *zap.Logger {
	conf := &appcontext.Config
	loggerObj, err := zlogger.New(conf.ServiceName, conf.LogEncoder,
		conf.LogLevel, conf.LogAddr, zap.Fields(
			zap.String("commit", buildCommit),
			zap.String("version", buildVersion),
			zap.String("build_date", buildDate),
		))
	fatalError(err, "init logger")

	// Register global logger
	zap.ReplaceGlobals(loggerObj)

	return loggerObj
}

func main() {
	conf := &appcontext.Config

	// Define cancelation context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	ctx = version.WithContext(ctx, &version.Version{
		Version: buildVersion,
		Commit:  buildCommit,
		Date:    buildDate,
	})

	// Init logger object
	loggerObj := initZapLogger()

	// Profiling server of collector
	profiler.Run(conf.Server.Profile.Mode,
		conf.Server.Profile.Listen, loggerObj, true)

	// Establish connect to the database
	fmt.Println("Connect to master database")
	masterDatabase, err := database.Connect(ctx,
		conf.System.Storage.MasterConnect, conf.IsDebug())
	fatalError(err, "connect to master database")

	fmt.Println("Connect to slave database")
	slaveDatabase, err := database.Connect(ctx,
		conf.System.Storage.SlaveConnect, conf.IsDebug())
	fatalError(err, "connect to slave database")

	// Register callback for history log
	fatalError(gormlog.Register(masterDatabase), "register history log callback")

	// Init permission manager
	permissionManager := permissions.NewManager(masterDatabase, conf.Permissions.RoleCacheLifetime)
	appinit.InitModelPermissions(permissionManager)

	// Init OAuth2 provider
	oauth2provider, jwtProvider := appinit.Auth(ctx, conf, masterDatabase)

	// Init messanger
	messangerObj := sendmsg.NewDefaultMessanger(emails.Templates())
	messangerObj.RegisterSender("log", wrapper.Sender(func(ctx context.Context, message sendmsg.Message) error {
		loggerObj.Info("Send message", zap.Any("message", message))
		return nil
	}))

	// Init email messanger if enabled
	if emCnf := &conf.Messanger.Email; emCnf.URL != "" && emCnf.APIKey != "" && emCnf.FromAddress != "" {
		email, err := email.New(email.WithConfig(emCnf.Mailer, &email.Config{
			URL:         emCnf.URL,
			APIKey:      emCnf.APIKey,
			Domain:      emCnf.Domain,
			FromAddress: emCnf.FromAddress,
			FromName:    emCnf.FromName,
			Password:    emCnf.Password,
			Port:        emCnf.Port,
		}), email.WithVars(map[string]any{
			"org": &conf.Messanger.EmailDefaults,
		}))
		fatalError(err, "init email messanger")
		messangerObj.RegisterSender("email", email)
	}

	messangerWrap := messangerWrapper(messangerObj)

	// Establish connection to Billing
	billingCl, err := billingClient.ConnectAPI(ctx, conf.Billing.Connect)
	fatalError(err, "connect to billing")
	defer func() {
		if err := billingCl.Close(); err != nil {
			loggerObj.Error("Close billing connection", zap.Error(err))
		}
	}()

	// Establish connection to Statistic
	statisticCl, err := statsClient.ConnectAPI(ctx, conf.Statistic.Connect)
	fatalError(err, "connect to statistic")

	rtbSourceUsecase := rtbsourceuc.New()

	statisticUsecase := statistic.NewUsecase(
		statistic.NewRepository(statisticCl))

	// Prepare context
	ctx = ctxlogger.WithLogger(ctx, loggerObj)
	ctx = database.WithDatabase(ctx, masterDatabase, slaveDatabase)
	ctx = permissions.WithManager(ctx, permissionManager)
	ctx = messanger.WithMessanger(ctx, messangerWrap)

	// Init HTTP server with GraphQL API
	httpServer := server.HTTPServer{
		SessionManager: appinit.SessionManager(conf.Session.CookieName, conf.Session.Lifetime),
		Authorizers: []auth.Authorizer{
			jwt.NewAuthorizer(jwtProvider),
			oauth2.NewAuthorizer(oauth2provider),
			devtoken.NewAuthorizer(gocast.IfThen(conf.IsDebug(), &devtoken.AuthOption{
				DevToken:     conf.Session.DevToken,
				DevUserID:    conf.Session.DevUserID,
				DevAccountID: conf.Session.DevAccountID,
			}, nil)),
		},
		ContextWrap: func(ctx context.Context) context.Context {
			ctx = ctxlogger.WithLogger(ctx, loggerObj)
			ctx = database.WithDatabase(ctx, masterDatabase, slaveDatabase)
			ctx = permissions.WithManager(ctx, permissionManager)
			ctx = messanger.WithMessanger(ctx, messangerWrap)
			return ctx
		},
		InitWrap: func(mux *chi.Mux) {
			// Register graphql playground with basic auth
			mux.With(basicauth.NewFromEnv("Graph", "GRAPHQL_USERS_")).
				Handle("/playground", playground.Handler("Query console", "/graphql"))

			// Init GraphQL API
			mux.Handle("/graphql", graphql.GraphQL(&resolvers.Usecases{
				Stats:     statisticUsecase,
				RTBSource: rtbSourceUsecase,
			}, jwtProvider))

			// Register OAuth2 providers
			if conf.SocialAuth.Facebook.IsValid() {
				oa2conf := conf.SocialAuth.Facebook.OAuth2Config("facebook")
				mux.Handle("/auth/facebook/*",
					rest.NewWrapper(facebook.NewFacebookConfig(oa2conf), rest.WithSessionProvider(jwtProvider)).
						HandleWrapper("/auth/facebook"),
				)
			}
		},
	}
	fatalError(httpServer.Run(ctx, conf.Server.HTTP.Listen), "HTTP server")
}

func fatalError(err error, msgs ...any) {
	if err != nil {
		log.Fatalln(append(msgs, err)...)
	}
}

func messangerWrapper(m sendmsg.Messanger) messanger.Messanger {
	return messanger.MessangerFunc(func(ctx context.Context, name string, recipients []string, vars map[string]any) error {
		return m.Send(ctx,
			sendmsg.WithTemplate(name),
			sendmsg.WithRecipients(recipients, nil, nil),
			sendmsg.WithVars(vars))
	})
}
