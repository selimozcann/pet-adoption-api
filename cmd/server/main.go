package main

import (
	"petapi/config"
	"petapi/internal/db"
	"petapi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db.Connect()

	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
