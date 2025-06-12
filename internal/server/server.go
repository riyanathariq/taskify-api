package server

import (
	"fmt"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
)

func Start() {
	fmt.Println(fmt.Sprintf("Starting and listening server on port %s...", dependencies.New().Config.AppPort))
}
