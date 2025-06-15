package repository

import (
	"github.com/riyanathariq/taskify-api/internal/repository/oauth_repo"
	"github.com/riyanathariq/taskify-api/internal/repository/user_repo"
	"gorm.io/gorm"
)

type Repositories struct {
	User  user_repo.UserRepository
	Oauth oauth_repo.OAuthRepository
}

func InitRepos(db *gorm.DB) *Repositories {
	return &Repositories{
		User:  user_repo.NewUserRepository(db),
		Oauth: oauth_repo.NewOAuthRepository(db),
	}
}
