package service

import (
	"errors"
	"time"

	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/badachirahul/dofocus-backend/internal/repository"
)

func StartFocusSession(taskID string, timerDurationSeconds int, userID string) (*models.FocusSession, error) {
	// Validate timer duration
	if timerDurationSeconds <= 0 {
		return nil, errors.New("invalid timer duration")
	}

	// Check if task exists and belongs to user
	task, err := repository.GetTaskByIDAndUserID(taskID, userID)

	if err != nil {
		return nil, errors.New("task not found")
	}

	// Check active session already exists
	activeSession, _ := repository.GetActiveSessionByUserID(userID)

	if activeSession != nil {
		return nil, errors.New("another active session already exists")
	}

	// Create session
	session, err := repository.CreateFocusSession(
		task.ID.String(),
		userID,
		timerDurationSeconds,
	)

	if err != nil {
		return nil, errors.New("failed to create focus session")
	}

	return session, nil
}

func PauseFocusSession(sessionID string, userID string) error {

	session, err := repository.GetActiveSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("active session not found")
	}

	// Calculate latest focused interval
	focusedDuration := int(
		time.Since(*session.LastResumedAt).Seconds(),
	)

	// Add into total focused seconds
	session.FocusedSeconds += focusedDuration

	// Update status
	session.Status = "paused"

	// Clear running timestamp
	session.LastResumedAt = nil

	return repository.UpdateFocusSession(session)
}

func ResumeFocusSession(sessionID string, userID string) error {

	session, err := repository.GetPausedSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("paused session not found")
	}

	now := time.Now()

	session.Status = "active"
	session.LastResumedAt = &now

	return repository.UpdateFocusSession(session)
}

func FinishFocusSession(sessionID string, userID string) error {

	session, err := repository.GetSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("session not found")
	}

	// If active, calculate latest interval
	if session.Status == "active" {

		focusedDuration := int(
			time.Since(*session.LastResumedAt).Seconds(),
		)

		session.FocusedSeconds += focusedDuration
	}

	now := time.Now()

	session.Status = "completed"
	session.EndedAt = &now
	session.LastResumedAt = nil

	// Save session
	err = repository.UpdateFocusSession(session)

	if err != nil {
		return err
	}

	// Mark task completed
	return repository.MarkTaskCompleted(
		session.TaskID.String(),
		userID,
	)
}

func CancelFocusSession(sessionID string, userID string) error {

	session, err := repository.GetSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("session not found")
	}

	// If active, calculate latest interval
	if session.Status == "active" {

		focusedDuration := int(
			time.Since(*session.LastResumedAt).Seconds(),
		)

		session.FocusedSeconds += focusedDuration
	}

	now := time.Now()

	session.Status = "cancelled"
	session.EndedAt = &now
	session.LastResumedAt = nil

	return repository.UpdateFocusSession(session)
}

func GetCurrentSession(taskID string,userID string) (*models.FocusSession, error) {

	session, err := repository.GetCurrentSessionByTaskIDAndUserID(
		taskID,
		userID,
	)

	if err != nil {
		return nil, errors.New("no active session found")
	}

	return session, nil
}