package archivarius

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// MonitoringUsecase wraps a Usecase to add monitoring metrics.
type MonitoringUsecase struct {
	usecase  Usecase
	counter  metric.Int64Counter
	duration metric.Int64Histogram
}

// NewMonitoringUsecase creates a new MonitoringUsecase with the provided metrics.
func NewMonitoringUsecase(usecase Usecase, counter metric.Int64Counter, duration metric.Int64Histogram) *MonitoringUsecase {
	return &MonitoringUsecase{
		usecase:  usecase,
		counter:  counter,
		duration: duration,
	}
}

// Statistic wraps the underlying usecase's Statistic method with metrics.
func (m *MonitoringUsecase) Statistic(ctx context.Context, opts ...Option) (*StatisticResponse, error) {
	start := time.Now()
	resp, err := m.usecase.Statistic(ctx, opts...)
	if err != nil {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Statistic"), attribute.String("error", err.Error())))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Statistic"), attribute.String("error", err.Error())))
	} else {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Statistic")))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Statistic")))
	}
	return resp, err
}

// Count wraps the underlying usecase's Count method with metrics.
func (m *MonitoringUsecase) Count(ctx context.Context, opts ...Option) (int64, error) {
	start := time.Now()
	resp, err := m.usecase.Count(ctx, opts...)
	if err != nil {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Count"), attribute.String("error", err.Error())))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Count"), attribute.String("error", err.Error())))
	} else {
		m.counter.Add(ctx, 1, metric.WithAttributes(attribute.String("method", "Count")))
		m.duration.Record(ctx, time.Since(start).Microseconds(), metric.WithAttributes(attribute.String("method", "Count")))
	}
	return resp, err
}
