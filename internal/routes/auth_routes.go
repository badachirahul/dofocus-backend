package routes

import (
	"github.com/badachirahul/dofocus-backend/internal/handler"
	"github.com/badachirahul/dofocus-backend/internal/middleware"
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

		taskRoutes := api.Group("/tasks")
		taskRoutes.Use(middleware.AuthMiddleware())
		{
			taskRoutes.POST("", handler.CreateTask)
			taskRoutes.GET("", handler.GetTasks)
			taskRoutes.PUT("/:id", handler.UpdateTask)
			taskRoutes.DELETE(("/:id"), handler.DeleteTask)
		}

		focusRoutes := api.Group("/focus")
		focusRoutes.Use(middleware.AuthMiddleware())
		{
			focusRoutes.POST("/start", handler.StartFocusSession)
			focusRoutes.POST("/pause/:sessionId", handler.PauseFocusSession)
			focusRoutes.POST("/resume/:sessionId", handler.ResumeFocusSession)
			focusRoutes.POST("/finish/:sessionId", handler.FinishFocusSession)
			focusRoutes.POST("/cancel/:sessionId", handler.CancelFocusSession)
			focusRoutes.GET("/task/:taskId", handler.GetCurrentSession)
		}
	}
}
