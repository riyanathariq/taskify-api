package handler_bank

import (
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
)

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "OK",
		"message": "Welcome to " + dependencies.New().Config.AppName,
	})
}
