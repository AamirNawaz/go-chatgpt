package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckAuth(c *gin.Context) {

	stringArray := c.Request.Header["Authorization"]
	justString := strings.Join(stringArray, " ")

	parsToken := strings.SplitAfter(justString, " ")

	if parsToken[0] == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Token is required",
		})
	}

	if parsToken[1] == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Token is required",
		})
	}
	_, err := jwt.Parse(parsToken[1], func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_ACCESS_TOKEN_SECRETE")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})

		c.Abort()
		return

	}
	c.Next()

}

func ExtractUserDetailsFromToken(c *gin.Context) string {
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
	}
	//fmt.Println("parse token:::", token)
	var userId string
	for key, val := range claims {
		// fmt.Printf("Key: %v, value: %v\n", key, val)
		if key == "iss" {
			userId = val.(string)
		}

	}
	return userId
}
