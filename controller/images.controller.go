package controller

import (
	"net/http"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/models"
	"github.com/gin-gonic/gin"
)

func CreateProductImage(c *gin.Context) {
	var productImage models.ProductImage
	c.BindJSON(&productImage)
	config.DB.Create(&productImage)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "ProductImage created successfully!", "resourceId": productImage.ID})
}

func FetchProductImage(c *gin.Context) {
	var productImage models.ProductImage
	productImageID := c.Param("product_image_id")
	config.DB.First(&productImage, productImageID)
	if productImage.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product image found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": productImage})
}

func UpdateProductImage(c *gin.Context) {
	var productImage models.ProductImage
	productImageID := c.Param("product_image_id")
	config.DB.First(&productImage, productImageID)
	if productImage.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product image found!"})
		return
	}
	c.BindJSON(&productImage)
	config.DB.Save(&productImage)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ProductImage updated successfully!"})
}

func DeleteProductImage(c *gin.Context) {
	var productImage models.ProductImage
	productImageID := c.Param("product_image_id")
	config.DB.First(&productImage, productImageID)
	if productImage.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No product image found!"})
		return
	}
	config.DB.Delete(&productImage)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "ProductImage deleted successfully!"})
}
