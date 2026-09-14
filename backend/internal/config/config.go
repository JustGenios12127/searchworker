package config

import (
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"time"
)

const (
	defaultHTTPAddr        = ":8080"
	defaultMigrationsDir = "migrations"
	defaultDatabasePingTimeout = 30 * time.Second
)


type Config struct {
	App     	AppConfig
	Auth      AuthConfig
	Chat      ChatConfig
	HTTP      HTTPConfig
	Database  DatabaseConfig
	Migration MigrationConfig
	Storage   StorageConfig

}

type AppConfig struct {
	PublicBaseURL      string
	PublicFilesBaseURL string
}

type AuthConfig struct {
	JWTSecret string
	JWTIssuer string
}

type ChatConfig struct {
	TelegramToken string
}

type HTTPConfig struct {
	Addr string
}

type DatabaseConfig struct {
	URL         string
	PingTimeout time.Duration
}

type MigrationConfig struct {
	Dir string
}

type StorageConfig struct {
	
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		App: AppConfig{
			PublicBaseURL:      getEnv("APP_PUBLIC_BASE_URL", ""),
			PublicFilesBaseURL: getEnv("APP_PUBLIC_FILES_BASE_URL", ""),
		},

		Auth: AuthConfig{
			JWTSecret: getEnv("AUTH_JWT_SECRET", ""),
			JWTIssuer: getEnv("AUTH_JWT_ISSUER", ""),
		},

		HTTP: HTTPConfig{
			Addr: getEnv("HTTP_ADDR", defaultHTTPAddr),
		},

		Database: DatabaseConfig{
			URL: getEnv("DATABASE_URL", ""),
			PingTimeout: getDurationEnv(
				"DATABASE_PING_TIMEOUT_SECONDS",
				defaultDatabasePingTimeout,
			),
		},

		Migration: MigrationConfig{
			Dir: getEnv("MIGRATIONS_DIR", defaultMigrationsDir),
		},
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getDurationEnv(key string, fallback time.Duration) time.Duration {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	seconds, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return time.Duration(seconds) * time.Second
}
