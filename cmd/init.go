package cmd

import (
	"fmt"
	"github.com/riyanathariq/taskify-api/database"
	"github.com/riyanathariq/taskify-api/internal/dependencies"
	"github.com/riyanathariq/taskify-api/internal/server"
	"log"
	"os"
)

func Start() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("[FATAL] Recovered from panic: %v\n", r)
		}
	}()

	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <command> <args>\n", os.Args[0])
		return
	}

	dependencies.LoadDependencies(
		dependencies.WithConfig(),
		dependencies.WithGormDB(),
		dependencies.WithRepository(),
	)

	switch os.Args[1] {
	case "api":
		server.Start()
	case "db:migrate":
		database.DatabaseMigration(dependencies.New().Config)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
