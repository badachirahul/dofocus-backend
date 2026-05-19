package routes

import (
	"fmt"

	"github.com/badachirahul/dofocus-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")
	{
		authRoutes := api.Group("/auth")
		{
			authRoutes.POST("/send-otp", handler.SendOTP)
			authRoutes.POST("/verify-otp", handler.VerifyOTP)
			authRoutes.POST("/register", handler.RegisterUser)
			authRoutes.POST("/login", handler.LoginUser)
		}

		taskRoutes := api.Group("/task")
		{
			fmt.Print(taskRoutes)
		}
	}
}
