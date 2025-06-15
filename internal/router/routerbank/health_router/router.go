package health_router

import (
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/handler/handler_bank"
)

func Router() func(r *gin.Engine) {
	return func(r *gin.Engine) {
		group := r.Group("/")
		group.GET("/ping", handler_bank.HealthCheck)
	}
}
