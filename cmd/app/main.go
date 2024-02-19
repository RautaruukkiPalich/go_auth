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
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}