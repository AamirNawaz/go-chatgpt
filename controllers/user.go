package controllers

import (
	db "go-chatgpt-app/config"
	"go-chatgpt-app/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db.DB.First(&users)

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   users,
	})
}

func GetUserProfile(c *gin.Context) {

	stringArray := c.Request.Header["Authorization"]
	justString := strings.Join(stringArray, " ")

	parsToken := strings.SplitAfter(justString, " ")

	if parsToken[0] == "" || parsToken[1] == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Token is required",
		})
	}

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(parsToken[1], claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRETE")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		c.Abort()
		return
	}
	//fmt.Println("parse token:::", token)
	var userId string
	for key, val := range claims {
		// fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "iss" {
			userId = val.(string)
		}
	}

	var user models.User
	db.DB.Where("id = ?", userId).First(&user)
	c.JSON(http.StatusUnauthorized, gin.H{
		"User": user,
	})

}
