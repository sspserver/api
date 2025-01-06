package runners

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/exporters/prometheus"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"

	// Internal project imports
	"github.com/geniusrabbit/archivarius/internal/archivarius"
	pb "github.com/geniusrabbit/archivarius/internal/archivarius/protobuf"
	"github.com/geniusrabbit/archivarius/internal/archivarius/repository"
	"github.com/geniusrabbit/archivarius/internal/config"
	"github.com/geniusrabbit/archivarius/internal/migratedb"
	grpcsrv "github.com/geniusrabbit/archivarius/internal/server/grpc"
)

// Archivarius initializes and runs the Archivarius service.
// It connects to the database, runs migrations if needed, sets up metrics,
// initializes repositories and use cases with monitoring, and starts the gRPC server.
func Archivarius(cfg *config.Config, logger *zap.Logger, runMigrations bool) error {
	var (
		err    error
		gormCl *gorm.DB
	)

	// Log the attempt to connect to Clickhouse
	logger.Info("Connecting to Clickhouse...")

	// Initialize GORM with Clickhouse driver and set logger to Info level
	gormCl, err = gorm.Open(clickhouse.Open(cfg.Clickhouse.DSN), &gorm.Config{
		Logger: gormlog.Default.LogMode(gormlog.Info),
	})
	if err != nil {
		logger.Error("Failed to connect to Clickhouse", zap.Error(err))
		return err
	}
	logger.Info("Successfully connected to Clickhouse")

	// Run database migrations if the runMigrations flag is set
	if runMigrations {
		logger.Info("Running database migrations...")
		err = migratedb.Migrate(cfg.Clickhouse.DSN, []migratedb.MigrateSource{
			{
				URI:                   []string{"file://deploy/migrations/stats"},
				SchemaMigrationsTable: "archivarius_schema_migration",
			},
		})
		if err != nil {
			logger.Error("Database migrations failed", zap.Error(err))
			return err
		}
		logger.Info("Database migrations completed successfully")
	}

	// Set up a context that listens for OS signals for graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// Initialize Prometheus exporter for OpenTelemetry metrics
	exporter, err := prometheus.New()
	if err != nil {
		logger.Error("Failed to create Prometheus exporter", zap.Error(err))
		return err
	}

	// Initialize OpenTelemetry metrics provider with the Prometheus exporter
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(exporter))
	meter := provider.Meter("billing")

	// Start serving metrics in a separate goroutine
	go serveMetrics(cfg.MetricsAddr, logger)

	// Initialize the repository with monitoring metrics
	logger.Info("Initializing repository with monitoring...")
	var repo archivarius.Repository
	repoCounter, err := repositoryRequestCountMetrics(meter)
	if err != nil {
		logger.Error("Failed to create repository request count metric", zap.Error(err))
		return err
	}
	repoDuration, err := repositoryRequestDurationMetrics(meter)
	if err != nil {
		logger.Error("Failed to create repository request duration metric", zap.Error(err))
		return err
	}
	repo = repository.NewRepository(gormCl)
	repo = repository.NewMonitoringRepository(repo, repoCounter, repoDuration)

	// Initialize the use case with monitoring metrics
	logger.Info("Initializing use case with monitoring...")
	var svc archivarius.Usecase
	usecaseCounter, err := usecaseRequestCountMetrics(meter)
	if err != nil {
		logger.Error("Failed to create usecase request count metric", zap.Error(err))
		return err
	}
	usecaseDuration, err := usecaseRequestDurationMetrics(meter)
	if err != nil {
		logger.Error("Failed to create usecase request duration metric", zap.Error(err))
		return err
	}
	svc = archivarius.NewUsecase(repo)
	svc = archivarius.NewMonitoringUsecase(svc, usecaseCounter, usecaseDuration)

	// Create an error group to manage goroutines with shared context
	g, ctx := errgroup.WithContext(ctx)

	// Start the gRPC server
	logger.Info("Starting gRPC server...")
	grpcServer := grpc.NewServer()
	grpcSrv := pb.NewGRPCServer(svc, cfg)
	grpcsrv.RegisterArchivariusServiceServer(grpcServer, grpcSrv)

	// Listen on the specified gRPC port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))
	if err != nil {
		logger.Error("Failed to listen on gRPC port", zap.Error(err))
		return err
	}

	// Add gRPC server to the error group
	g.Go(func() error {
		logger.Info("gRPC server listening", zap.String("address", lis.Addr().String()))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("gRPC server encountered an error", zap.Error(err))
			return fmt.Errorf("failed to serve gRPC: %w", err)
		}
		logger.Info("gRPC server stopped")
		return nil
	})

	// Add a goroutine to handle graceful shutdown on context cancellation
	g.Go(func() error {
		<-ctx.Done()
		logger.Info("Shutting down gRPC server gracefully...")
		grpcServer.GracefulStop()
		return nil
	})

	// Wait for all goroutines to finish
	err = g.Wait()
	cancel() // Ensure context is canceled
	if err != nil {
		logger.Error("Service encountered an error", zap.Error(err))
	} else {
		logger.Info("Service shutdown gracefully")
	}

	return err
}

// serveMetrics starts an HTTP server that serves Prometheus metrics at the /metrics endpoint.
func serveMetrics(addr string, logger *zap.Logger) {
	logger.Info("Serving Prometheus metrics", zap.String("address", addr))
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(addr, nil); err != nil {
		logger.Error("Error serving metrics", zap.Error(err))
	}
}
