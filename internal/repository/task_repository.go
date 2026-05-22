package repository

import (
	"errors"

	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/google/uuid"
)

func CreateTask(taskName string, userID string) (*models.Task, error) {

	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		return nil, err
	}

	task := models.Task{
		TaskName: taskName,
		UserID:   parsedUserID,
	}

	err = database.DB.Create(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
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

func UpdateTask(taskID string, taskName string, completed bool, userID string) error {
	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error

	if err != nil {
		return errors.New("task not found")
	}

	task.TaskName = taskName
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

	// Delete all focus sessions related to this task and user
	err = database.DB.
		Where("task_id = ? AND user_id = ?", taskID, userID).
		Delete(&models.FocusSession{}).Error

	if err != nil {
		return err
	}

	// Delete task
	return database.DB.Delete(&task).Error
}

func GetTask(taskID string, userID string) (*models.Task, error) {
	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}