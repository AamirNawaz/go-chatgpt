package controllers

import (
	"fmt"
	Database "go-chatgpt-app/config"
	dto "go-chatgpt-app/dto"
	"go-chatgpt-app/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var JwtSecreteKey = []byte("myjwtsecret")

func Login(c *gin.Context) {
	var userInput models.User

	validation.Errors{
		"email":    validation.Validate(userInput.Email, validation.Required, is.Email),
		"password": validation.Validate(userInput.Password, validation.Required, validation.Length(4, 12)),
	}.Filter()

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//Now fetch user
	var userResult models.User
	Database.DB.Where("email=?", userInput.Email).First(&userResult)
	if userResult.Id == 0 {
		c.JSON(500, gin.H{
			"success": false,
			"message": "Email not found!",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(userResult.Password, []byte(userInput.Password))
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": "incorrect password!",
		})
		return
	}

	//generate jwt token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   userResult.Email,
		Issuer:    strconv.Itoa(int(userResult.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	})

	tokenString, err := claims.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRETE")))
	fmt.Print(tokenString)
	if err != nil {
		c.JSON(403, gin.H{
			"status":  "error",
			"message": "Token Expired or invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func Signup(c *gin.Context) {
	var userInput dto.UserInputDto
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	validation.Errors{
		"name":     validation.Validate(userInput.Name, validation.Required),
		"email":    validation.Validate(userInput.Email, validation.Required, is.Email),
		"password": validation.Validate(userInput.Password, validation.Required),
	}.Filter()

	// Validate the parsed fields
	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	user := models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: hashedPassword,
		Status:   "active",
	}

	// Insert the user into the database
	if err := Database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//creating user response which will not send password to api response
	userResponse := dto.SignupUserResponse{
		ID:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Status: user.Status,
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "User created successfully",
		"Data":    userResponse,
	})
}
