package controller

import (
	"net/http"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/models"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Category
	c.BindJSON(&category)
	config.DB.Create(&category)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Category created successfully!", "resourceId": category.ID})
}

func FetchCategory(c *gin.Context) {
	var category models.Category
	categoryID := c.Param("category_id")
	config.DB.First(&category, categoryID)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No category found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": category})
}

func UpdateCategory(c *gin.Context) {
	var category models.Category
	categoryID := c.Param("category_id")
	config.DB.First(&category, categoryID)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No category found!"})
		return
	}
	c.BindJSON(&category)
	config.DB.Save(&category)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Category updated successfully!"})
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	categoryID := c.Param("category_id")
	config.DB.First(&category, categoryID)
	if category.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No category found!"})
		return
	}
	config.DB.Delete(&category)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Category deleted successfully!"})
}
