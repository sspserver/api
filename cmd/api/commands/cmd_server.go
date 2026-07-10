package commands

import (
	"context"
	"strings"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/demdxx/gocast/v2"
	"github.com/demdxx/sendmsg"
	"github.com/demdxx/sendmsg/sender/email"
	"github.com/demdxx/sendmsg/sender/wrapper"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/geniusrabbit/blaze-api/pkg/appcmd"
	blazeauth "github.com/geniusrabbit/blaze-api/pkg/auth"
	"github.com/geniusrabbit/blaze-api/pkg/auth/elogin/facebook"
	"github.com/geniusrabbit/blaze-api/pkg/auth/jwt"
	"github.com/geniusrabbit/blaze-api/pkg/auth/oauth2"
	"github.com/geniusrabbit/blaze-api/pkg/context/ctxlogger"
	"github.com/geniusrabbit/blaze-api/pkg/database"
	"github.com/geniusrabbit/blaze-api/pkg/messanger"
	"github.com/geniusrabbit/blaze-api/pkg/permissions"
	"github.com/geniusrabbit/blaze-api/pkg/profiler"
	"github.com/geniusrabbit/blaze-api/repository/account"
	accAuth "github.com/geniusrabbit/blaze-api/repository/account/auth"
	accountauth "github.com/geniusrabbit/blaze-api/repository/account/authorizer"
	accountgraphql "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql"
	accountlogin "github.com/geniusrabbit/blaze-api/repository/account/delivery/graphql/account_login"
	accountrepo "github.com/geniusrabbit/blaze-api/repository/account/repository"
	accountuc "github.com/geniusrabbit/blaze-api/repository/account/usecase"
	"github.com/geniusrabbit/blaze-api/repository/historylog/middleware/gormlog"
	optionrp "github.com/geniusrabbit/blaze-api/repository/option/repository"
	optionuc "github.com/geniusrabbit/blaze-api/repository/option/usecase"
	rbacrepo "github.com/geniusrabbit/blaze-api/repository/rbac/repository"
	"github.com/geniusrabbit/blaze-api/repository/socialauth/delivery/rest"
	socialauthrepo "github.com/geniusrabbit/blaze-api/repository/socialauth/repository"
	socialauthuc "github.com/geniusrabbit/blaze-api/repository/socialauth/usecase"

	"github.com/sspserver/api/cmd/api/appcontext"
	"github.com/sspserver/api/cmd/api/appinit"
	"github.com/sspserver/api/cmd/api/server"
	"github.com/sspserver/api/pkg/models"
	rtbsourceuc "github.com/sspserver/api/pkg/repository/rtbsource/usecase"
	statisticrc "github.com/sspserver/api/pkg/repository/statistic/repository"
	statisticuc "github.com/sspserver/api/pkg/repository/statistic/usecase"
	"github.com/sspserver/api/pkg/server/graphql"
	gqlaccounts "github.com/sspserver/api/pkg/server/graphql/accounts"
	"github.com/sspserver/api/pkg/server/graphql/resolvers"
	"github.com/sspserver/api/pkg/server/graphql/wiring"
	"github.com/sspserver/api/pkg/sysops"
	"github.com/sspserver/api/pkg/user"
	"github.com/sspserver/api/private/emails"
)

// APICommand is the main API server command
var APICommand = &appcmd.Command[appcontext.ConfigType]{
	Name:     "server",
	HelpDesc: "Run API server",
	Exec:     apiCommand,
}

func apiCommand(ctx context.Context, _ []string, conf *appcontext.ConfigType) error {
	loggerObj := ctxlogger.Get(ctx)

	// Profiling server or collector
	profiler.Run(conf.Server.Profile.Mode,
		conf.Server.Profile.Listen, loggerObj, true)

	// Establish connect to the database
	masterDatabase, err := database.Connect(ctx, conf.System.Storage.MasterConnect)
	if err != nil {
		return err
	}

	slaveDatabase, err := database.Connect(ctx, conf.System.Storage.SlaveConnect)
	if err != nil {
		return err
	}

	// Register callback for history log
	if err := gormlog.Register(masterDatabase); err != nil {
		return err
	}

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

	// Init email sender if configured
	if emCnf := &conf.Messanger.Email; emCnf.URL != "" && emCnf.APIKey != "" && emCnf.FromAddress != "" {
		emailSender, err := email.New(email.WithConfig(emCnf.Mailer, &email.Config{
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
		if err != nil {
			return err
		}
		messangerObj.RegisterSender("email", emailSender)
	}

	messangerWrap := messangerWrapper(messangerObj)

	// Establish connection to Statistic database
	statDatabase, err := database.Connect(ctx, conf.System.Statistic.Connect)
	if err != nil {
		return err
	}

	// Init statistic usecase
	statisticUsecase := statisticuc.NewUsecase(
		statisticrc.NewRepository(statDatabase))

	// Init RTB Source usecase
	rtbSourceUsecase := rtbsourceuc.New()

	// Init Options usecase
	optionsUsecase := optionuc.NewUsecase(optionrp.NewOptionRepository(map[string]any{
		"ad.rtb.domain": conf.Options.RTBServerDomain,
		"ad.template.code": prepareAdCode(conf.Options.AdTemplateCode,
			conf.Options.JSSDKDomain, conf.Options.RTBServerDomain),
		"ad.direct.url": prepareAdCode(conf.Options.AdDirectTemplateURL,
			conf.Options.JSSDKDomain, conf.Options.RTBServerDomain),
		"ad.direct.code": prepareAdCode(conf.Options.AdDirectTemplateCode,
			conf.Options.JSSDKDomain, conf.Options.RTBServerDomain),
	}))

	// Init system options
	sysops.Set(`system.hostname`, conf.Hostname)
	sysops.Set(`system.datacenter`, conf.DatacenterName)
	sysops.Set(`logic.crud.default.approval`, true)

	// Prepare context
	ctx = ctxlogger.WithLogger(ctx, loggerObj)
	ctx = database.WithDatabase(ctx, masterDatabase, slaveDatabase)
	ctx = permissions.WithManager(ctx, permissionManager)
	ctx = messanger.WithMessanger(ctx, messangerWrap)

	userModule := user.NewModule(func() *models.User { return &models.User{} })
	accountRepoInst := accountrepo.NewSessionRepository(
		func() *models.User { return &models.User{} },
		func() *models.Account { return &models.Account{} },
		func() *account.Member[*models.User, *models.Account] {
			return &account.Member[*models.User, *models.Account]{}
		},
	)
	memberRepoInst := accountrepo.NewMemberRepositoryFor(
		func() *account.Member[*models.User, *models.Account] {
			return &account.Member[*models.User, *models.Account]{}
		},
	)
	accountUC := accountuc.NewAccountUsecase(userModule.Repo, accountRepoInst, memberRepoInst)
	memberUC := accountuc.NewMemberUsecase(userModule.Repo, accountRepoInst, memberRepoInst)
	socialAuthUsecase := socialauthuc.New(socialauthrepo.New(), userModule.Repo)
	authLoader := accAuth.NewLoader(userModule.Repo, accountRepoInst, memberRepoInst)

	// Ensure superuser exists
	if err = appinit.EnsureSuperuser(
		ctx,
		conf.Superuser.Email,
		conf.Superuser.Password,
		userModule.Repo,
		accountRepoInst,
		memberRepoInst,
	); err != nil {
		return err
	}

	// Init HTTP server with GraphQL API
	httpServer := server.HTTPServer{
		SessionManager: appinit.SessionManager(conf.Session.CookieName, conf.Session.Lifetime),
		AuthLoader:     authLoader,
		Authorizers: []blazeauth.Authorizer[*models.User, *models.Account]{
			jwt.NewAuthorizer(jwtProvider, authLoader),
			oauth2.NewAuthorizer(oauth2provider, accountRepoInst),
			accountauth.NewDevTokenAuthorizer(gocast.IfThen(conf.IsDebug(), &accountauth.AuthOption{
				DevToken:     conf.Session.DevToken,
				DevUserID:    conf.Session.DevUserID,
				DevAccountID: conf.Session.DevAccountID,
			}, nil), authLoader),
		},
		ContextWrap: func(ctx context.Context) context.Context {
			ctx = ctxlogger.WithLogger(ctx, loggerObj)
			ctx = database.WithDatabase(ctx, masterDatabase, slaveDatabase)
			ctx = permissions.WithManager(ctx, permissionManager)
			ctx = messanger.WithMessanger(ctx, messangerWrap)
			return ctx
		},
		GraphqlOptions: graphql.Options{
			graphql.WithUserAccountResolvers(
				jwtProvider,
				wiring.NewUserQueryResolver(userModule.Core, userModule.Email, userModule.Password),
				accountgraphql.NewAuthResolver(
					jwtProvider,
					accountRepoInst,
					accountUC,
					rbacrepo.New(),
				),
				wiring.NewAccountQueryResolver(
					wiring.AccountQueryResolverConfig{
						Users:    userModule.Core,
						Accounts: accountUC,
						Members:  memberUC,
					},
				),
				accountgraphql.NewMemberQueryResolver(
					accountgraphql.MemberQueryResolverConfig[*models.User, *models.Account, *gqlaccounts.Account]{
						Accounts: accountUC,
						Members:  memberUC,
						UserRepo: userModule.Repo,
					},
				),
			),
			graphql.WithUserLoginHandler(
				jwtProvider,
				accountlogin.NewEmailPasswordLogin(userModule.Email, userModule.Password.Repo()),
				accountRepoInst,
			),
		},
		InitWrap: func(mux *chi.Mux) {
			// Register graphql playground
			mux.Handle("/playground", playground.Handler("Query console", "/graphql"))

			// Init GraphQL API
			mux.Handle("/graphql", graphql.GraphQL(&resolvers.Usecases{
				Stats:     statisticUsecase,
				RTBSource: rtbSourceUsecase,
				Options:   optionsUsecase,
			}, jwtProvider))

			// Register OAuth2 providers
			if conf.SocialAuth.Facebook.IsValid() {
				oa2conf := conf.SocialAuth.Facebook.OAuth2Config("facebook")
				mux.Handle("/auth/facebook/*",
					rest.NewWrapper(
						facebook.NewFacebookConfig(oa2conf),
						rest.WithSessionProvider(jwtProvider),
						rest.WithSocialAuthUsecase(socialAuthUsecase),
						rest.WithAccountResolver(func(ctx context.Context, filter *account.Filter) ([]*models.Account, error) {
							return accountRepoInst.FetchList(ctx, filter)
						}),
					).
						HandleWrapper("/auth/facebook"),
				)
			}
		},
	}
	return httpServer.Run(ctx, conf.Server.HTTP.Listen)
}

func messangerWrapper(m sendmsg.Messanger) messanger.Messanger {
	return messanger.MessangerFunc(func(ctx context.Context, name string, recipients []string, vars map[string]any) error {
		return m.Send(ctx,
			sendmsg.WithTemplate(name),
			sendmsg.WithRecipients(recipients, nil, nil),
			sendmsg.WithVars(vars))
	})
}

func prepareAdCode(templateCode, jssdkDomain, adServerDomain string) string {
	return strings.NewReplacer(
		"\\n", "\n",
		"{JSSDK_DOMAIN}", jssdkDomain,
		"{ADSERVER_DOMAIN}", adServerDomain,
	).Replace(templateCode)
}
