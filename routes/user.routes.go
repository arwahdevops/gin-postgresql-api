package routes

import (
	"github.com/ArwahDevops/gin-postgresql-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user", controller.GetUsers)
	router.POST("/user", controller.CreateUser)
	router.DELETE("/user/:id", controller.DeleteUser)
	router.PUT("/user/:id", controller.UpdateUser)
}
