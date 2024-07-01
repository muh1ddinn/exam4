package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	DebugMode   = "debug"
	TestMode    = "test"
	ReleaseMode = "release"
)

// Config holds the configuration for the client
type Config struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresPassword string
	PostgresUser     string
	PostgresDatabase string
	UserServicePort  string
	UserServiceHost  string
	OrderServiceHost string
	OrderServicePort string
}

// Load loads environment variables and inflates Config
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("error !!!", err)
	}
	cfg := Config{}

	cfg.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", DebugMode))

	cfg.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	cfg.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "manager_admin"))
	cfg.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "muhiddin"))
	cfg.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1"))

	cfg.OrderServiceHost = cast.ToString(getOrReturnDefault("SERVICE_NAME", "CRM_Admin Service"))
	cfg.OrderServicePort = cast.ToString(getOrReturnDefault("CONTENT_GRPC_PORT", ":8082"))

	cfg.UserServicePort = "8082"
	cfg.UserServiceHost = "localhost"

	return cfg
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
