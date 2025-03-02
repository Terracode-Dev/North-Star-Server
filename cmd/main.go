package main

import (
	"github.com/Terracode-Dev/North-Star-Server/internal/server"

	_ "github.com/Terracode-Dev/North-Star-Server/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @host localhost:3000
// @BasePath /hr-api
func main() {
	server.InitServer()
}
