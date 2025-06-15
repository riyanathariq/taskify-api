package router

import (
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/router/routerbank/auth_router"
	"github.com/riyanathariq/taskify-api/internal/router/routerbank/health_router"
)

func RegisterRouter(r *gin.Engine) {
	InitRouter(r,
		health_router.Router(),
		auth_router.Router(),
	)
}

type RouteInitializer func(*gin.Engine)

func InitRouter(r *gin.Engine, initializers ...RouteInitializer) {
	for _, init := range initializers {
		init(r)
	}
}
