package config

import (
	"os"
	"strconv"
)

type Config struct {
	// JWT config
	JWT_SECRET_KEY string
	JWT_ISSUER     string
	JWT_EXPIRE     int64
}

var config *Config = nil

func getEnvDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func stringToNumber(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)

	if err == nil {
		return number
	}

	panic(err)
}

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			JWT_SECRET_KEY: getEnvDefault("JWT_SECRET_KEY", "secret"),
			JWT_ISSUER:     getEnvDefault("JWT_ISSUER", "go-jwt-auth-project"),
			JWT_EXPIRE:     stringToNumber(getEnvDefault("JWT_EXPIRE", "3600")),
		}
	}

	return config
}
