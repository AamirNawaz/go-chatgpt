package controllers

import (
	db "go-chatgpt-app/config"
	"go-chatgpt-app/models"
	"net/http"

	middleware "go-chatgpt-app/middleware"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
	})
}

func GetUserProfile(c *gin.Context) {
	//extracting userId from token
	userId := middleware.ExtractUserDetailsFromToken(c)

	var user models.User
	db.DB.Where("id = ?", userId).First(&user)
	c.JSON(http.StatusUnauthorized, gin.H{
		"User": user,
	})

}
