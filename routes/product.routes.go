package routes

import (
	"github.com/ArwahDevops/gin-postgresql-api/controller"
	"github.com/ArwahDevops/gin-postgresql-api/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.POST("/products/create", middleware.AuthMiddleware, controller.CreateProduct)
	router.GET("/products/find", middleware.AuthMiddleware, controller.FetchProduct)
	router.PUT("/products/:id", middleware.AuthMiddleware, controller.UpdateProduct)
	router.DELETE("/products/:id", middleware.AuthMiddleware, controller.DeleteProduct)
}

func CategoryRoutes(router *gin.Engine) {
	router.POST("/category/create", middleware.AuthMiddleware, controller.CreateCategory)
	router.GET("/category/find", middleware.AuthMiddleware, controller.FetchCategory)
	router.PUT("/category/:id", middleware.AuthMiddleware, controller.UpdateCategory)
	router.DELETE("/category/:id", middleware.AuthMiddleware, controller.DeleteCategory)
}
