package server

import "github.com/rautaruukkipalich/go_auth/internal/store"

type Config struct {
	BindAddress string
	LogLevel    string
	Store       *store.Config
}


func NewConfig() *Config {
	return &Config{
		BindAddress: "localhost:8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}