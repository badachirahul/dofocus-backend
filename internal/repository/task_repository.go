package repository

import (
	"errors"

	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/google/uuid"
)

func CreateTask(title string, userID string) error {
	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		return err
	}

	task := models.Task{
		Title:  title,
		UserID: parsedUserID,
	}

	return database.DB.Create(&task).Error
}

func GetTasks(userID string) ([]models.Task, error) {
	var tasks []models.Task

	err := database.DB.
		Where("user_id = ?", userID).
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(taskID string, title string, completed bool, userID string) error {
	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error

	if err != nil {
		return errors.New("task not found")
	}

	task.Title = title
	task.Completed = completed

	return database.DB.Save(&task).Error
}

func DeleteTask(taskID string, userID string) error {
	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error

	if err != nil {
		return errors.New("task not found")
	}

	return database.DB.Delete(&task).Error
}
