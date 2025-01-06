package runners

import "go.opentelemetry.io/otel/metric"

func repositoryRequestCountMetrics(meter metric.Meter) (metric.Int64Counter, error) {
	return meter.Int64Counter(
		"repository_requests_count",
		metric.WithDescription("Repository requests count"),
	)
}

func repositoryRequestDurationMetrics(meter metric.Meter) (metric.Int64Histogram, error) {
	return meter.Int64Histogram(
		"repository_request_duration_microseconds",
		metric.WithUnit("microseconds"),
		metric.WithDescription("Repository request latency"))
}

func usecaseRequestCountMetrics(meter metric.Meter) (metric.Int64Counter, error) {
	return meter.Int64Counter(
		"usecase_requests_count",
		metric.WithDescription("Usecase requests count"),
	)
}

func usecaseRequestDurationMetrics(meter metric.Meter) (metric.Int64Histogram, error) {
	return meter.Int64Histogram(
		"usecase_request_duration_microseconds",
		metric.WithUnit("microseconds"),
		metric.WithDescription("Usecase request latency"))
}
