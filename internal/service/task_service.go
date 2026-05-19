package service

import (
	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/badachirahul/dofocus-backend/internal/repository"
)

func CreateTask(title string, userID string) error {
	return repository.CreateTask(title, userID)
}

func GetTasks(userID string) ([]models.Task, error) {
	return repository.GetTasks(userID)
}

func UpdateTask(taskID string, title string, completed bool, userID string) error {
	return repository.UpdateTask(
		taskID,
		title,
		completed,
		userID,
	)
}

func DeleteTask(taskID string, userID string) error {
	return repository.DeleteTask(taskID, userID)
}