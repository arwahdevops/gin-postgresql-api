package controller

import (
	"net/http"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data user tidak ditemukan",
		})
		return
	}
	config.DB.Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data user tidak ditemukan",
		})
		return
	}
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
