package store

type Config struct {
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		DatabaseURL: "postgres://postgres:postgres@localhost:5432/notes?sslmode=disable",
	}
}