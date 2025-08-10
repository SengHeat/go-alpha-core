package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv  string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret        string
	JWTExpiryMinutes int
}

func Load() *Config {
	_ = godotenv.Load()

	config := &Config{
		AppEnv:           getEnv("APP_ENV", "development"),
		AppPort:          getEnv("APP_PORT", "8080"),
		DBHost:           getEnv("DB_HOST", "localhost"),
		DBPort:           getEnv("DB_PORT", "5432"),
		DBUser:           getEnv("DB_USER", "postgres"),
		DBPass:           getEnv("DB_PASS", "password"),
		DBName:           getEnv("DB_NAME", "myapp"),
		JWTSecret:        getEnv("JWT_SECRET", "mysecret"),
		JWTExpiryMinutes: getEnvAsInt("JWT_EXPIRY_MINUTES", 60),
	}

	return config
}

func getEnv(key, defaultVal string) string {
	env := os.Getenv(key)

	if env == "" {
		return defaultVal
	}

	return env
}

func getEnvAsInt(key string, defaultVal int) int {
	env := os.Getenv(key)

	if env == "" {
		return defaultVal
	}

	value, err := strconv.Atoi(env)
	if err != nil {
		return defaultVal
	}

	return value
}
