package handler

import (
	"fmt"
	"net/http"

	"github.com/badachirahul/dofocus-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type StartFocusSessionRequest struct {
	TaskID               string `json:"task_id" binding:"required"`
	TimerDurationSeconds int    `json:"timer_duration_seconds" binding:"required"`
}

func StartFocusSession(c *gin.Context) {

	var request StartFocusSessionRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	fmt.Println("Frontend request: ", request)

	userID := c.GetString("user_id")

	session, err := service.StartFocusSession(
		request.TaskID,
		request.TimerDurationSeconds,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Focus session started",
		"session": session,
	})
}

func PauseFocusSession(c *gin.Context) {

	sessionID := c.Param("sessionId")

	userID := c.GetString("user_id")

	err := service.PauseFocusSession(
		sessionID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session paused",
	})
}

func ResumeFocusSession(c *gin.Context) {

	sessionID := c.Param("sessionId")

	userID := c.GetString("user_id")

	err := service.ResumeFocusSession(
		sessionID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session resumed",
	})
}

func FinishFocusSession(c *gin.Context) {

	sessionID := c.Param("sessionId")

	userID := c.GetString("user_id")

	err := service.FinishFocusSession(
		sessionID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task completed successfully",
	})
}

func CancelFocusSession(c *gin.Context) {

	sessionID := c.Param("sessionId")

	userID := c.GetString("user_id")

	err := service.CancelFocusSession(
		sessionID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session cancelled",
	})
}

func GetCurrentSession(c *gin.Context) {

	taskID := c.Param("taskId")

	userID := c.GetString("user_id")

	session, err := service.GetCurrentSession(
		taskID,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "No active session found",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session found",
		"data":    session,
	})
}
