package handler

import (
	"fmt"
	"net/http"

	"github.com/badachirahul/dofocus-backend/internal/service"
	"github.com/gin-gonic/gin"
)

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func SendOTP(c *gin.Context) {
	var request SendOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email format",
		})
		return
	}

	err := service.CheckUserAlreadyExists(request.Email)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}

	otp, err := service.GenerateAndSaveOTP(request.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save OTP",
		})
		return
	}

	fmt.Print(otp)

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP generated successfully",
	})
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

func VerifyOTP(c *gin.Context) {
	var request VerifyOTPRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	err := service.VerifyUserOTP(request.Email, request.OTP)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OTP verified successfully",
	})
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func RegisterUser(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	err := service.RegisterUser(
		request.Name,
		request.Email,
		request.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginUser(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	token, err := service.LoginUser(
		request.Email,
		request.Password,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}