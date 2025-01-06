package config

// Clickhouse configuration
type Clickhouse struct {
	DSN string `json:"dsn" yaml:"dsn" env:"CLICKHOUSE_DSN"`
}

// Config is a service configuration
type Config struct {
	ServiceName string `json:"service_name" yaml:"service_name" env:"SERVICE_NAME"`

	LivenessAddr  string `json:"liveness_addr" yaml:"liveness_addr" env:"LIVENESS_ADDR"`
	ReadinessAddr string `json:"readiness_addr" yaml:"readiness_addr" env:"READINESS_ADDR"`
	MetricsAddr   string `json:"metrics_addr" yaml:"metrics_addr" env:"METRICS_ADDR"`
	GRPCPort      string `json:"grpc_port" yaml:"grpc_port" env:"GRPC_PORT"`

	LogLevel   string `json:"log_level" yaml:"log_level" env:"LOG_LEVEL"`
	LogEncoder string `json:"log_encoder" yaml:"log_encoder" env:"LOG_ENCODER"`
	LogAddr    string `json:"log_addr" yaml:"log_addr" env:"LOG_ADDR"`

	Clickhouse Clickhouse `json:"clickhouse" yaml:"clickhouse"`
}

func (cfg *Config) IsDebug() bool {
	return cfg.LogLevel == "debug"
}
