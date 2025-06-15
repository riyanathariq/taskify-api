package user_repo

import (
	"context"
	"github.com/riyanathariq/taskify-api/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, payload *models.User) error
	DetailByID(ctx context.Context, id string) (*models.User, error)
	DetailByUsername(ctx context.Context, username string) (*models.User, error)
	List(ctx context.Context, payload *models.User) ([]*models.User, error)
	Update(ctx context.Context, id string, update *models.User) error
	Delete(ctx context.Context, id string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}
