package auth_router

import (
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/handler/handler_bank/auth_handler"
	"github.com/riyanathariq/taskify-api/internal/middlewares"
)

func Router() func(r *gin.Engine) {
	return func(r *gin.Engine) {
		group := r.Group("/api/v1/auth")
		LogoutGroup := r.Group("/api/v1/auth")

		LogoutGroup.Use(middlewares.AuthMiddleware())

		group.POST("/login", authhandler.LoginHandler)
		LogoutGroup.POST("/logout", authhandler.LogoutHandler)
	}
}
