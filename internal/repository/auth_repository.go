package repository

import (
	"time"

	"errors"

	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/models"
)

func UserExists(email string) (bool, error) {
	var user models.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err == nil {
		return true, nil
	}

	return false, nil
}

func SaveOrUpdateOTP(email string, otp string) error {
	var otpVerification models.OTPVerification

	err := database.DB.
		Where("email = ?", email).
		First(&otpVerification).Error

	if err == nil {

		otpVerification.OTP = otp
		otpVerification.Verified = false
		otpVerification.ExpiresAt = time.Now().Add(5 * time.Minute)

		return database.DB.Save(&otpVerification).Error
	}

	newOTP := models.OTPVerification{
		Email:     email,
		OTP:       otp,
		Verified:  false,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}

	return database.DB.Create(&newOTP).Error
}

func VerifyOTP(email string, otp string) error {
	var otpVerification models.OTPVerification

	err := database.DB.
		Where("email = ?", email).
		First(&otpVerification).Error

	if err != nil {
		return errors.New("OTP not found")
	}

	if otpVerification.OTP != otp {
		return errors.New("Invalid OTP")
	}

	if time.Now().After(otpVerification.ExpiresAt) {
		return errors.New("OTP expired")
	}

	otpVerification.Verified = true

	return database.DB.Save(&otpVerification).Error
}

func IsEmailVerified(email string) (bool, error) {
	var otpVerification models.OTPVerification

	err := database.DB.
		Where("email = ?", email).
		First(&otpVerification).Error

	if err != nil {
		return false, err
	}

	return otpVerification.Verified, nil
}

func CreateUser(name string, email string, password string) error {
	user := models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return database.DB.Create(&user).Error
}

func DeleteOTPVerification(email string) error {
	return database.DB.
		Where("email = ?", email).
		Delete(&models.OTPVerification{}).Error
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}