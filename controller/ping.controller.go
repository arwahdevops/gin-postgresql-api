package controller

import "github.com/gin-gonic/gin"

func PingController(c *gin.Context) {
	c.String(200, "pong")
}
