package routes

import (
	"github.com/ArwahDevops/gin-postgresql-api/controller"
	"github.com/gin-gonic/gin"
)

func PingRoutes(router *gin.Engine) {
	router.GET("/ping", controller.PingController)
}
