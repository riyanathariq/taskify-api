package dependencies

import (
	"github.com/riyanathariq/taskify-api/internal/config"
	"github.com/riyanathariq/taskify-api/internal/repository"
	"github.com/riyanathariq/taskify-api/pkg/gorm"
)

type Options func(*Dependency)

func WithConfig() Options {
	return func(dep *Dependency) {
		dep.Config = config.LoadConfig()
	}
}

func WithGormDB() Options {
	return func(d *Dependency) {
		_, _ = gorm.InitGormDB(d.Config)
	}
}

func WithRepository() Options {
	return func(d *Dependency) {
		d.Repository = repository.InitRepos(d.GormDB)
	}
}
