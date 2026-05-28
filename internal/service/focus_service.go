package service

import (
	"errors"
	"time"

	"github.com/badachirahul/dofocus-backend/internal/models"
	"github.com/badachirahul/dofocus-backend/internal/repository"
)

func StartFocusSession(taskID string, timerDurationSeconds int, userID string) (*models.FocusSession, error) {

	if timerDurationSeconds <= 0 {
		return nil, errors.New("invalid timer duration")
	}

	task, err := repository.GetTaskByIDAndUserID(taskID, userID)

	if err != nil {
		return nil, errors.New("task not found")
	}

	activeSession, _ := repository.GetActiveSessionByUserID(userID)

	if activeSession != nil {
		return nil, errors.New("another active session already exists")
	}

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

func FinishFocusSession(sessionID string, userID string) error {

	session, err := repository.GetSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("session not found")
	}

	if session.Status == "active" {

		session.FocusedSeconds =
			int(time.Since(session.StartedAt).Seconds())

		if session.FocusedSeconds >=
			session.TimerDurationSeconds {

			session.FocusedSeconds =
				session.TimerDurationSeconds
		}
	}

	now := time.Now()

	session.Status = "completed"
	session.EndedAt = &now

	return repository.UpdateFocusSession(session)
}

func CancelFocusSession(sessionID string, userID string) error {

	session, err := repository.GetSessionByIDAndUserID(
		sessionID,
		userID,
	)

	if err != nil {
		return errors.New("session not found")
	}

	if session.Status == "active" {

		session.FocusedSeconds =
			int(time.Since(session.StartedAt).Seconds())

		if session.FocusedSeconds >=
			session.TimerDurationSeconds {

			session.FocusedSeconds =
				session.TimerDurationSeconds
		}
	}

	now := time.Now()

	session.Status = "cancelled"
	session.EndedAt = &now

	return repository.UpdateFocusSession(session)
}

func GetCurrentSession(
	taskID string,
	userID string,
) (*models.FocusSession, error) {

	session, err := repository.GetCurrentSessionByTaskIDAndUserID(
		taskID,
		userID,
	)

	if err != nil {
		return nil, errors.New("no active session found")
	}

	if session.Status == "active" {

		session.FocusedSeconds =
			int(time.Since(session.StartedAt).Seconds())

		if session.FocusedSeconds >=
			session.TimerDurationSeconds {

			session.FocusedSeconds =
				session.TimerDurationSeconds
		}
	}

	return session, nil
}