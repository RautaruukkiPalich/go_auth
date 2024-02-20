package config

import (
	"fmt"
	"os"
	"strconv"
)


var (
	DB_USER = GetEnv("DB_USER", "postgres")
	DB_PASS = GetEnv("DB_PASS", "postgres")
	DB_NAME = GetEnv("DB_NAME", "go_auth_users")
	DB_HOST = GetEnv("DB_HOST", "localhost")
	DB_PORT = GetEnv("DB_PORT", "5432")

	HOST_ADDR = GetEnv("HOST_ADDR", "localhost")
	HOST_PORT = GetEnv("HOST_PORT", "8080")
	LOG_LEVEL = GetEnv("LOG_LEVEL", "info")

	JWT_SECRET_KEY  = GetEnv("JWT_SECRET_KEY", "secretkey")
	JWT_TTL_SECONDS = GetEnvAsInt("JWT_TTL_SECONDS", 3600)

	BIND_ADDR = fmt.Sprintf("%s:%s", HOST_ADDR, HOST_PORT)
	DB_URI    = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
)


func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvAsInt(name string, defaultVal int) int {
	valueStr := GetEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
