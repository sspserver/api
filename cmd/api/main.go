package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/demdxx/goconfig"
	"go.uber.org/zap"

	"github.com/geniusrabbit/blaze-api/pkg/appcmd"
	"github.com/geniusrabbit/blaze-api/pkg/context/version"
	_ "github.com/geniusrabbit/blaze-api/pkg/gopentracing"
	"github.com/geniusrabbit/blaze-api/pkg/migratedb"
	"github.com/geniusrabbit/blaze-api/pkg/zlogger"

	"github.com/sspserver/api/cmd/api/appcontext"
	"github.com/sspserver/api/cmd/api/commands"
	"github.com/sspserver/api/pkg/sysops"
)

var (
	buildDate    = ""
	buildCommit  = ""
	buildVersion = "develop"
)

func init() {
	fmt.Println()
	fmt.Println("███████ ███████ ██████         █████  ██████  ██")
	fmt.Println("██      ██      ██   ██       ██   ██ ██   ██ ██")
	fmt.Println("███████ ███████ ██████  █████ ███████ ██████  ██")
	fmt.Println("     ██      ██ ██            ██   ██ ██      ██")
	fmt.Println("███████ ███████ ██            ██   ██ ██      ██")
	fmt.Println()
	fmt.Println("Version:", buildVersion, " (", buildCommit, ")")
	fmt.Println("Build date:", buildDate)
	fmt.Println()

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
		fatalError(migratedb.Migrate(context.Background(), conf.System.Storage.MasterConnect, []migratedb.MigrateSource{
			{
				URI:                   []string{"file:///data/migrations/initial"},
				SchemaMigrationsTable: "schema_migrations_initial",
			},
			{
				URI:                   []string{"file:///data/migrations/project"},
				SchemaMigrationsTable: "schema_migrations",
			},
			{
				URI:                   []string{"file:///data/migrations/fixtures"},
				SchemaMigrationsTable: "schema_migrations_fixtures",
			},
		}), "migrate database")
	}
}

func main() {
	conf := &appcontext.Config

	// Define cancelation context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Init logger for startup
	loggerObj, err := zlogger.New(conf.ServiceName, conf.LogEncoder,
		conf.LogLevel, conf.LogAddr, zap.Fields(
			zap.String("commit", buildCommit),
			zap.String("version", buildVersion),
			zap.String("build_date", buildDate),
		))
	fatalError(err, "init logger")
	zap.ReplaceGlobals(loggerObj)

	// Set build-time system options
	sysops.Set(`system.version`, buildVersion)
	sysops.Set(`system.commit`, buildCommit)
	sysops.Set(`system.build_date`, buildDate)

	// Application with command list
	app := &appcmd.App{
		Name:        "api",
		Description: "SSP API - Supply Side Platform API Server",
		Version:     buildVersion,
		BuildCommit: buildCommit,
		BuildDate:   buildDate,
		CmdList: appcmd.ICommands{
			commands.APICommand,
		},
		BeforeCommandRun: func(ctx context.Context, cmd appcmd.ICommand) (context.Context, error) {
			fmt.Println()
			fmt.Println("░█ Log Level:\x1b[32m", conf.LogLevel, "\x1b[0m")
			fmt.Println("░█ Run command:\x1b[31m", cmd.Cmd(), "\x1b[0m")

			// Register version information
			ctx = version.WithContext(ctx, &version.Version{
				Version: buildVersion,
				Commit:  buildCommit,
				Date:    buildDate,
			})

			return ctx, nil
		},
	}

	fatalError(app.Run(ctx, os.Args), "application run")
}

func fatalError(err error, msgs ...any) {
	if err != nil {
		zap.L().Fatal(fmt.Sprint(msgs...), zap.Error(err))
	}
}
