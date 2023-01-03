package routes

import (
	"github.com/ArwahDevops/gin-postgresql-api/controller"
	"github.com/ArwahDevops/gin-postgresql-api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/login", controller.UserLogin)
	router.POST("/register", controller.RegisterUser)
	router.DELETE("/user/:id", middleware.AuthMiddleware, controller.DeleteUser)
	router.PUT("/user/:id", middleware.AuthMiddleware, controller.UpdateUser)
}
