package main

import (
	"go-chatgpt-app/config"
	"go-chatgpt-app/routes"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	app := gin.Default()
	//Database connection
	config.Connect()

	//handling cors
	app.Use(cors.Default())

	//routes
	routes.RoutesSetup(app)

	app.Run("localhost:9080")

}
