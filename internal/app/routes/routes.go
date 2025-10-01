package routes

import (
	"GoNotebookRealtime/internal/app/controllers/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	apiRoutes := router.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			// authRoutes.POST("/login", gin.WrapF(auth.Login))
			authRoutes.POST("/register", gin.WrapF(auth.Register))
		}
	}

	// Handle 404 - Not Found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error": "Page not found",
		})
	})

	// userRoutes := router.Group("/user")
	// {
	// 	userRoutes.GET("/:id", user.GetUser)
	// 	userRoutes.PUT("/:id", user.UpdateUser)
	// }
}
