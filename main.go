package main

import (
	"os"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Running Mode
	config.RunMode()
	// Database
	config.ConnectDB()
	router := gin.Default()

	routes.PingRoutes(router)
	routes.UserRoutes(router)

	appPort := os.Getenv("APP_PORT")
	router.Run(":" + appPort)

}
