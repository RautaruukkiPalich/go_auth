package server

import "github.com/rautaruukkipalich/go_auth/config"

type Config struct {
	BindAddress string
	LogLevel    string
	DatabaseURL string
}

func NewConfig() *Config {
	// return &Config{
	// 	BindAddress: "localhost:8080",
	// 	LogLevel: "debug",
	// 	DatabaseURL: "postgres://postgres:postgres@localhost:5432/go_auth_users?sslmode=disable",
	// }
	return &Config{
		BindAddress: config.BIND_ADDR,
		LogLevel:    config.LOG_LEVEL,
		DatabaseURL: config.DB_URI,
	}
}
