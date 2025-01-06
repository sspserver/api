package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/alecthomas/kong"
	"github.com/demdxx/goconfig"
	"github.com/geniusrabbit/blaze-api/pkg/zlogger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/geniusrabbit/archivarius/cmd/runners"
	"github.com/geniusrabbit/archivarius/internal/config"
)

// Build information variables.
// These are typically set during the build process.
var (
	buildDate    = ""
	buildCommit  = ""
	buildVersion = "develop"
)

// initZapLogger initializes a Zap logger based on the provided configuration.
// It includes build metadata such as commit, version, and build date.
func initZapLogger(cfg *config.Config) *zap.Logger {
	loggerObj, err := zlogger.New(cfg.ServiceName, cfg.LogEncoder,
		cfg.LogLevel, cfg.LogAddr, zap.Fields(
			zap.String("commit", buildCommit),
			zap.String("version", buildVersion),
			zap.String("build_date", buildDate),
		))
	if err != nil {
		log.Fatalln(err.Error())
	}

	return loggerObj
}

// cli defines the command-line interface structure using Kong.
// It includes options for memory profiling, configuration file, and archivarius commands.
var cli struct {
	Memprofile  string `cmd:"" help:"Path to memory profile file"`
	Config      string `help:"Config file" short:"c"`
	Archivarius struct {
		RunMigrations bool `help:"Run Migrations"`
	} `cmd:"" help:"Archivarius"`
}

// run is the main execution function.
// It parses command-line arguments, loads configuration, initializes logging,
// handles memory profiling if requested, and executes the appropriate command.
func run() error {
	// Parse command-line arguments using Kong.
	ctx := kong.Parse(&cli,
		kong.Name("archivarius"),
		kong.Description("Archivarius"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))

	// Load configuration with defaults, environment variables, and config file.
	cfg := &config.Config{}
	err := goconfig.Load(cfg, goconfig.WithDefaults(), goconfig.WithEnv(), goconfig.WithFile(cli.Config))
	if err != nil {
		return err
	}

	// Initialize the Zap logger.
	logger := initZapLogger(cfg)

	if cfg.IsDebug() {
		logger.Info("Config", zap.Any("config", cfg))
	}

	// Handle memory profiling if the Memprofile flag is set.
	if cli.Memprofile != "" {
		f, err := os.Create(cli.Memprofile)
		if err != nil {
			logger.Fatal(fmt.Sprintf("profile: could not create memory profile %q: %v", cli.Memprofile, err))
		}

		// Save the old memory profile rate and set a new one.
		old := runtime.MemProfileRate
		runtime.MemProfileRate = 4096
		logger.Info(fmt.Sprintf("profile: memory profiling enabled (rate %d), %s", runtime.MemProfileRate, cli.Memprofile))

		// Ensure that the memory profile is written and resources are cleaned up on exit.
		defer func() {
			_ = pprof.Lookup("heap").WriteTo(f, 0)
			f.Close()
			runtime.MemProfileRate = old
			logger.Info(fmt.Sprintf("profile: memory profiling disabled, %s", cli.Memprofile))
		}()
	}

	// Execute the appropriate command based on the parsed context.
	switch ctx.Command() {
	case "archivarius":
		return runners.Archivarius(
			cfg,
			logger,
			cli.Archivarius.RunMigrations,
		)
	}
	return nil
}

// main is the entry point of the application.
// It calls the run function and logs any errors that occur.
func main() {
	err := run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
