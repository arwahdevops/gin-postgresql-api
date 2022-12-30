package routes

import (
	"github.com/ArwahDevops/gin-postgresql-api/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/login", controller.UserLogin)
	router.POST("/register", controller.RegisterUser)
	router.DELETE("/user/:id", controller.AuthMiddleware, controller.DeleteUser)
	router.PUT("/user/:id", controller.AuthMiddleware, controller.UpdateUser)
}
