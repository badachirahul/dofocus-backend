package handler

import (
	"net/http"

	"github.com/badachirahul/dofocus-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type CreateTaskRequest struct {
	Title string `json:"title" binding:"required"`
}

func CreateTask(c *gin.Context) {
	var request CreateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userID := c.GetString("user_id")

	err := service.CreateTask(request.Title, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
	})
}

func GetTasks(c *gin.Context) {
	userID := c.GetString("user_id")

	tasks, err := service.GetTasks(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

type UpdateTaskRequest struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func UpdateTask(c *gin.Context) {
	taskID := c.Param("id")

	userID := c.GetString("user_id")

	var request UpdateTaskRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	err := service.UpdateTask(
		taskID,
		request.Title,
		request.Completed,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
	})
}

func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")

	userID := c.GetString("user_id")

	err := service.DeleteTask(taskID, userID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
	})
}