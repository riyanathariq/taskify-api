package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
	"github.com/riyanathariq/taskify-api/internal/router"
)

func Start() {
	r := gin.Default()

	router.RegisterRouter(r)

	_ = r.Run(fmt.Sprintf(":%s", dependencies.New().Config.AppPort))
}
