package routes

import (
	"go-chatgpt-app/controllers"

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
	// user := router.Group("/api/users")
	// {
	// 	user.GET("/", controllers.GetUsers)
	// 	user.DELETE("/delete/:id", middleware.CheckAuth, middleware.CheckRole, controllers.DeleteUser)
	// }

	//Api Routes
	// appApi := router.Group("/api/search")
	// {
	// 	appApi.POST("/search", controllers.PromptSearch)
	// 	appApi.GET("/search-response", controllers.PromptResponse)
	// 	appApi.DELETE("/delete-prompt/:id", controllers.PromptResponse)
	// }

}
