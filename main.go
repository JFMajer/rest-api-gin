package main

import (
	"github.com/JFMajer/rest-api-gin/db"
	"github.com/JFMajer/rest-api-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
