package repository

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/geniusrabbit/archivarius/internal/archivarius"
)

// MonitoringRepository wraps a archivarius.Repository to add monitoring metrics.
type MonitoringRepository struct {
	Repository archivarius.Repository
	counter    metric.Int64Counter
	duration   metric.Int64Histogram
}

// NewMonitoringRepository creates a new MonitoringRepository with the provided metrics.
func NewMonitoringRepository(Repository archivarius.Repository, counter metric.Int64Counter, duration metric.Int64Histogram) *MonitoringRepository {
	return &MonitoringRepository{
		Repository: Repository,
		counter:    counter,
		duration:   duration,
	}
}

// Search wraps the underlying repository's Statistic method with metrics.
func (m *MonitoringRepository) Statistic(ctx context.Context, opts ...archivarius.Option) ([]*archivarius.AdItem, error) {
	start := time.Now()
	resp, err := m.Repository.Statistic(ctx, opts...)
	if err != nil {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Statistic"), attribute.String("error", err.Error())))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Statistic"), attribute.String("error", err.Error())))
	} else {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Statistic")))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Statistic")))
	}
	return resp, err
}

// Count wraps the underlying repository's Count method with metrics.
func (m *MonitoringRepository) Count(ctx context.Context, opts ...archivarius.Option) (int64, error) {
	start := time.Now()
	resp, err := m.Repository.Count(ctx, opts...)
	if err != nil {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Count"), attribute.String("error", err.Error())))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Count"), attribute.String("error", err.Error())))
	} else {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Count")))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Count")))
	}
	return resp, err
}
