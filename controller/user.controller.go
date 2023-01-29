package controller

import (
	"net/http"

	"github.com/ArwahDevops/gin-postgresql-api/config"
	"github.com/ArwahDevops/gin-postgresql-api/middleware"
	"github.com/ArwahDevops/gin-postgresql-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func RegisterUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	var existingUser models.User
	config.DB.Where("email = ?", user.Email).First(&existingUser)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email sudah terdaftar",
		})
		return
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengenkripsi password pengguna",
		})
		return
	}
	user.Password = hashedPassword

	config.DB.Create(&user)
	c.JSON(200, &user)
}

func UserLogin(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email atau password tidak boleh kosong",
		})
		return
	}
	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}
	if !checkPassword(user.Password, request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Email atau password salah",
		})
		return
	}
	token, err := middleware.GenerateJWT(uint64(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal membuat JWT",
		})
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(200, gin.H{
		"message": "Login sukses",
		"jwt":     token,
	})
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

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengenkripsi password pengguna",
		})
		return
	}
	user.Password = hashedPassword

	config.DB.Save(&user)
	c.JSON(200, &user)
}
