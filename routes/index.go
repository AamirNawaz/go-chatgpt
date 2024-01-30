package routes

import (
	"go-chatgpt-app/ApiClient"
	"go-chatgpt-app/controllers"
	"go-chatgpt-app/middleware"

	"github.com/gin-gonic/gin"
)

func RoutesSetup(router *gin.Engine) {

	//Auth routes
	auth := router.Group("/api/auth")
	{
		auth.POST("/signup", controllers.Signup)
		auth.POST("/login", controllers.Login)

		// auth.POST("/get-token", controllers.GetNewAccessToken)
		// auth.GET("/logout", controllers.Logout)
	}

	//User Routes
	user := router.Group("/api/users")
	{
		user.GET("/", middleware.CheckAuth, controllers.GetUsers)
		user.GET("/profile", middleware.CheckAuth, controllers.GetUserProfile)
		// user.DELETE("/delete/:id", middleware.CheckAuth, middleware.CheckRole, controllers.DeleteUser)
	}

	//Api Routes
	appApi := router.Group("/api/search")
	{
		appApi.GET("/prompt-search", ApiClient.HttpApiRequest)
		//appApi.GET("/search-response", controllers.PromptResponse)
	}

}
