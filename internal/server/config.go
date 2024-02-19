package server

type Config struct {
	BindAddress string
	LogLevel    string
	DatabaseURL string
}


func NewConfig() *Config {
	return &Config{
		BindAddress: "localhost:8080",
		LogLevel: "debug",
		DatabaseURL: "postgres://postgres:postgres@localhost:5432/go_auth_users?sslmode=disable",
	}
}