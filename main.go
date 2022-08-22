package main

import (
	"log"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env file")
	}
}

func main() {
	router := gin.New()
	config.ConnectDB()

	routes.PingRoutes(router)
	routes.UserRoutes(router)

	router.Run(":8080")

}
