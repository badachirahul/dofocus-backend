package service

import (
	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/badachirahul/dofocus-backend/internal/repository"
)

func CreateTask(taskName string, userID string) (*models.Task, error) {
	return repository.CreateTask(taskName, userID)
}

func GetTasks(userID string) ([]models.Task, error) {
	return repository.GetTasks(userID)
}

func UpdateTask(taskID string, taskName string, completed bool, userID string) error {
	return repository.UpdateTask(
		taskID,
		taskName,
		completed,
		userID,
	)
}

func DeleteTask(taskID string, userID string) error {
	return repository.DeleteTask(taskID, userID)
}