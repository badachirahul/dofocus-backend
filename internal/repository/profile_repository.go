package repository

import (
	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/models"
)

type HeatmapData struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

func GetProfile(userID string) (map[string]interface{}, error) {
	var user models.User

	// Fetch user
	err := database.DB.
		Where("user_id = ?", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	var heatmapData []HeatmapData

	// Aggregate focus data
	err = database.DB.
		Table("focus_sessions").
		Select(`
		TO_CHAR(
			started_at AT TIME ZONE 'Asia/Kolkata',
			'YYYY-MM-DD'
		) as date,
		SUM(focused_seconds) as count
	`).
		Where("user_id = ? AND status = ?", userID, "completed").
		Group("TO_CHAR(started_at AT TIME ZONE 'Asia/Kolkata', 'YYYY-MM-DD')").
		Order("TO_CHAR(started_at AT TIME ZONE 'Asia/Kolkata', 'YYYY-MM-DD')").
		Scan(&heatmapData).Error

	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"user": map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
			"year":  user.CreatedAt.Year(),
		},
		"heatmapData": heatmapData,
	}
	return response, nil
}

type TaskFocusData struct {
	TaskName       string `json:"taskName"`
	FocusedSeconds int    `json:"focusedSeconds"`
}

func GetDailyProfile(userID string, date string) (map[string]interface{}, error) {
	var hourlyFocus int

	// Total daily focused seconds
	err := database.DB.
		Table("focus_sessions").
		Select("COALESCE(SUM(focused_seconds), 0)").
		Where(`
			focus_sessions.user_id = ?
			AND focus_sessions.status = ?
			AND DATE(
				focus_sessions.started_at AT TIME ZONE 'Asia/Kolkata'
			) = ?
		`, userID, "completed", date).
		Scan(&hourlyFocus).Error

	if err != nil {
		return nil, err
	}

	var tasks []TaskFocusData

	// Task-wise focus aggregation
	err = database.DB.
		Table("focus_sessions").
		Select(`
			tasks.task_name as task_name,
			SUM(focus_sessions.focused_seconds) as focused_seconds
		`).
		Joins(`
			JOIN tasks
			ON tasks.id = focus_sessions.task_id
		`).
		Where(`
			focus_sessions.user_id = ?
			AND focus_sessions.status = ?
			AND DATE(
				focus_sessions.started_at AT TIME ZONE 'Asia/Kolkata'
			) = ?
		`, userID, "completed", date).
		Group("tasks.task_name").
		Order("focused_seconds DESC").
		Scan(&tasks).Error

	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{
		"hourlyFocus": hourlyFocus,
		"tasks":       tasks,
	}

	return response, nil
}
