package dependencies

import (
	"github.com/riyanathariq/taskify-api/internal/config"
	"github.com/riyanathariq/taskify-api/internal/repository"
	"github.com/riyanathariq/taskify-api/pkg/gorm"
	"log"
)

type Options func(*Dependency)

func WithConfig() Options {
	return func(dep *Dependency) {
		dep.Config = config.LoadConfig()
	}
}

func WithGormDB() Options {
	return func(d *Dependency) {
		gormDB, err := gorm.InitGormDB(d.Config)
		if err != nil {
			log.Fatalf("failed to init gorm db: %v", err)
		}

		d.GormDB = gormDB
	}
}

func WithRepository() Options {
	return func(d *Dependency) {
		d.Repository = repository.InitRepos(d.GormDB)
	}
}
