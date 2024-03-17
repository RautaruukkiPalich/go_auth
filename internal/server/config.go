package server

import (
	"fmt"
	"net/http"

	"github.com/rautaruukkipalich/go_auth/pkg/env"
)

type (
	Config struct {
		BindAddress string
		LogLevel    string
		DatabaseURL string
	}

	CorsConfig struct {
		AllowedOrigins []string
		AllowedMethods []string
		AllowedHeaders []string
		AllowCredentials bool
	}
)

func NewConfig() *Config {
	LOG_LEVEL := env.GetEnv("LOG_LEVEL", "debug")

	BIND_ADDR := fmt.Sprintf(
		"%s:%s",
		env.GetEnv("HOST_ADDR", "127.0.0.1"),
		env.GetEnv("HOST_PORT", "8081"),
	)

	DB_URI := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		env.GetEnv("DB_USER", "postgres"),
		env.GetEnv("DB_PASS", "postgres"),
		env.GetEnv("DB_HOST", "localhost"),
		env.GetEnv("DB_PORT", "5432"),
		env.GetEnv("DB_NAME", "postgres"),
	)

	return &Config{
		BindAddress: BIND_ADDR,
		LogLevel:    LOG_LEVEL,
		DatabaseURL: DB_URI,
	}
}

func NewCorsConfig() *CorsConfig {
	ORIGINS := []string{
		"http://127.0.0.1:8080",
		"http://localhost:8080",
		"127.0.0.1:8080",
		"localhost:8080",
	}
	
	METHODS := []string{
		http.MethodGet, 
		http.MethodPost, 
		http.MethodPatch,
	}

	HEADERS := []string{
		"Authorization",
		"Accept",
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Credentials",
		"Content-Type",
		"Content-Length",
		"X-Requested-With",
	}
	
	CREDENTIALS := true

	return &CorsConfig{
		AllowedOrigins: ORIGINS,
		AllowedMethods: METHODS,
		AllowedHeaders: HEADERS,
		AllowCredentials: CREDENTIALS,
	}
	
}
