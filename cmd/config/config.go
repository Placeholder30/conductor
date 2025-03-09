package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lpernett/godotenv"
)

type PgConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

var Envs = initConfig()

func (config PgConfig) FormatDSN() string {
	return fmt.Sprintf("host='%s' port=%d user='%s' "+
		"password='%s' dbname='%s' sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)
}

func initConfig() PgConfig {
	godotenv.Load()

	return PgConfig{
		Host:     getEnv("PG_HOST", "localhost"),
		Port:     getEnvAsInt("PG_PORT", 5432),
		User:     getEnv("PG_USER", "postgres"),
		Password: getEnv("PG_PASSWORD", "mysecretpassword"),
		Dbname:   getEnv("PG_DB_NAME", "conductor"),
	}
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
