package config

import (
	"os"
)

type Config struct {
	ServiceName  string
	OTLPEndpoint string
	Environment  string
}

func Load() Config {
	return Config{
		ServiceName:  getenv("SERVICE_NAME", "wallet-service"),
		OTLPEndpoint: getenv("OTEL_EXPORTER_OTLP_ENDPOINT", "http://otel-collector:4318"),
		Environment:  getenv("ENV", "dev"),
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
