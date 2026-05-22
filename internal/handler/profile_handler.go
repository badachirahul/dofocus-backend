package handler

import (
	"fmt"
	"net/http"

	"github.com/badachirahul/dofocus-backend/internal/service"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	requestedUserID := c.Param("userId")

	loggedInUserID := c.GetString("user_id")

	// Security validation
	if requestedUserID != loggedInUserID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized access",
		})
		return
	}

	profileData, err := service.GetProfile(requestedUserID)

	fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch profile",
		})
		return
	}

	c.JSON(http.StatusOK, profileData)
}

func GetDailyProfile(c *gin.Context) {
	requestedUserID := c.Param("userId")

	date := c.Param("date")

	loggedInUserID := c.GetString("user_id")

	// Security validation
	if requestedUserID != loggedInUserID {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Unauthorized access",
		})
		return
	}

	profileData, err := service.GetDailyProfile(
		requestedUserID,
		date,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch daily profile",
		})
		return
	}

	c.JSON(http.StatusOK, profileData)
}