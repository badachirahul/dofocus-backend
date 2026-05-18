package routes

import (
	"github.com/badachirahul/dofocus-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DoFocus backend running",
		})
	})

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/send-otp", handler.SendOTP)
		authRoutes.POST("/verify-otp", handler.VerifyOTP)

		authRoutes.POST("/register", handler.RegisterUser)
		authRoutes.POST("/login", handler.LoginUser)
	}
}