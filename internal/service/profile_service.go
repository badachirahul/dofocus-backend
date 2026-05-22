package service

import "github.com/badachirahul/dofocus-backend/internal/repository"

func GetProfile(userID string) (map[string]interface{}, error) {
	return repository.GetProfile(userID)
}

func GetDailyProfile(userID string, date string) (map[string]interface{}, error) {
	return repository.GetDailyProfile(
		userID,
		date,
	)
}