package server

import "github.com/rs/cors"

func newCors(cfg *CorsConfig) *cors.Cors {
	return cors.New(
		cors.Options{
			AllowedOrigins:   cfg.AllowedOrigins,
			AllowedMethods:   cfg.AllowedMethods,
			AllowedHeaders:   cfg.AllowedHeaders,
			AllowCredentials: cfg.AllowCredentials,
		},
	)
}