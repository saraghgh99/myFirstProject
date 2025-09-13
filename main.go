package main

import (
	"rest-api-pra/db"
	"rest-api-pra/routes"
    "github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
