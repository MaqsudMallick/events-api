package main

import (
	"log"
	"module-name/db"
	"module-name/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(`Starting server...`)
	db.Init()
	server := gin.Default()
	routes.InitRoutes(server)
	server.Run(":8080")
}
