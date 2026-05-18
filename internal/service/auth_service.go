package service

import (
	"errors"
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"

	"github.com/badachirahul/dofocus-backend/internal/repository"
	"github.com/badachirahul/dofocus-backend/internal/util"
)

func CheckUserAlreadyExists(email string) error {
	exists, err := repository.UserExists(email)

	if err != nil {
		return err
	}

	if exists {
		return errors.New("email already registered")
	}

	return nil
}

func GenerateAndSaveOTP(email string) (string, error) {
	otp := fmt.Sprintf("%06d", rand.Intn(1000000))

	err := repository.SaveOrUpdateOTP(email, otp)

	if err != nil {
		return "", err
	}

	err = util.SendOTPEmail(email, otp)

	if err != nil {
		return "", err
	}

	return otp, nil
}

func VerifyUserOTP(email string, otp string) error {
	return repository.VerifyOTP(email, otp)
}

func RegisterUser(name string, email string, password string) error {
	verified, err := repository.IsEmailVerified(email)

	if err != nil {
		return err
	}

	if !verified {
		return errors.New("email not verified")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	err = repository.CreateUser(
		name,
		email,
		string(hashedPassword),
	)

	if err != nil {
		return err
	}

	return repository.DeleteOTPVerification(email)
}

func LoginUser(email string, password string) (string, error) {
	user, err := repository.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	token, err := util.GenerateJWT(user.UserID.String())

	if err != nil {
		return "", err
	}

	return token, nil
}
