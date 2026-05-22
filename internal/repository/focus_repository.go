package repository

import (
	"time"

	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/google/uuid"
)

func GetTaskByIDAndUserID(taskID string, userID string) (*models.Task, error) {

	var task models.Task

	err := database.DB.
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func GetActiveSessionByUserID(userID string) (*models.FocusSession, error) {

	var session models.FocusSession

	err := database.DB.
		Where("user_id = ? AND status = ?", userID, "active").
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func CreateFocusSession(taskID string, userID string, timerDurationSeconds int) (*models.FocusSession, error) {

	parsedTaskID, err := uuid.Parse(taskID)

	if err != nil {
		return nil, err
	}

	parsedUserID, err := uuid.Parse(userID)

	if err != nil {
		return nil, err
	}

	now := time.Now()

	session := models.FocusSession{
		SessionID:             uuid.New(),
		TaskID:                parsedTaskID,
		UserID:                parsedUserID,
		TimerDurationSeconds:  timerDurationSeconds,
		FocusedSeconds:        0,
		Status:                "active",
		LastResumedAt:         &now,
		StartedAt:             now,
	}

	err = database.DB.Create(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func GetActiveSessionByIDAndUserID(sessionID string, userID string) (*models.FocusSession, error) {

	var session models.FocusSession

	err := database.DB.
		Where(
			"session_id = ? AND user_id = ? AND status = ?",
			sessionID,
			userID,
			"active",
		).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func UpdateFocusSession(
	session *models.FocusSession,
) error {

	return database.DB.Save(session).Error
}

func GetPausedSessionByIDAndUserID(sessionID string, userID string) (*models.FocusSession, error) {

	var session models.FocusSession

	err := database.DB.
		Where(
			"session_id = ? AND user_id = ? AND status = ?",
			sessionID,
			userID,
			"paused",
		).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func GetSessionByIDAndUserID(sessionID string, userID string) (*models.FocusSession, error) {

	var session models.FocusSession

	err := database.DB.
		Where(
			"session_id = ? AND user_id = ?",
			sessionID,
			userID,
		).
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}

func MarkTaskCompleted(taskID string, userID string) error {

	var task models.Task

	err := database.DB.
		Where(
			"id = ? AND user_id = ?",
			taskID,
			userID,
		).
		First(&task).Error

	if err != nil {
		return err
	}

	task.Completed = true

	return database.DB.Save(&task).Error
}

func GetCurrentSessionByTaskIDAndUserID(taskID string, userID string) (*models.FocusSession, error) {

	var session models.FocusSession

	err := database.DB.
		Where(
			"task_id = ? AND user_id = ? AND status IN ?",
			taskID,
			userID,
			[]string{"active", "paused"},
		).
		Order("created_at desc").
		First(&session).Error

	if err != nil {
		return nil, err
	}

	return &session, nil
}