package main

import (
	"log"
	"github.com/rautaruukkipalich/go_auth/internal/server"
)

//	@title			Swagger Example API
//	@version		0.0.1
//	@description	This is a sample Auth service

//	@host		localhost:8080
//	@BasePath	/

//	@SecurityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	config := server.NewConfig()
	if err := server.Start(config); err != nil {
		log.Fatal(err)
	}
}

// func init() {
//     // loads values from .env into the system
//     log.Print("try load env...")
// 	if err := godotenv.Load(); err != nil {
//         log.Print("No .env file found")
// 		return
//     }
// 	log.Print("ENV file loaded")
// }
