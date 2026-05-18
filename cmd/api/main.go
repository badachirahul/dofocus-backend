package main

import (
	"github.com/gin-gonic/gin"

	"github.com/badachirahul/dofocus-backend/internal/database"
	"github.com/badachirahul/dofocus-backend/internal/routes"
)

func main() {

	database.ConnectDatabase()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
