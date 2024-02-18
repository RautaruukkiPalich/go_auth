package main

import (
	"log"
	"github.com/rautaruukkipalich/go_auth/internal/server"
)


//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Notes
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/
func main() {
	config := server.NewConfig()

	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}