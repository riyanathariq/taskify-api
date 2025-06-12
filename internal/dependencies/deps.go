package dependencies

import (
	"github.com/riyanathariq/taskify-api/internal/config"
	"github.com/riyanathariq/taskify-api/internal/repository"
	"gorm.io/gorm"
)

type Dependency struct {
	Config     *config.Config
	GormDB     *gorm.DB
	Repository *repository.Repositories
}
