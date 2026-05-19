package routes

import (
	"fmt"

	"github.com/badachirahul/dofocus-backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	api := router.Group("/api/v1")
	{
		authRoutes1 := api.Group("/auth")
		{
			authRoutes1.POST("/send-otp", handler.SendOTP)
			authRoutes1.POST("/verify-otp", handler.VerifyOTP)
			authRoutes1.POST("/register", handler.RegisterUser)
			authRoutes1.POST("/login", handler.LoginUser)
		}

		authRoutes2 := api.Group("/task")
		{
			fmt.Print(authRoutes2)
		}
	}
}
