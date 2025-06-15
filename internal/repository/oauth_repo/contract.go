package oauth_repo

import (
	"context"
	"github.com/riyanathariq/taskify-api/internal/models"
	"gorm.io/gorm"
)

type OAuthRepository interface {
	SaveToken(ctx context.Context, token *models.Token) error
	FindTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error)
	FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*models.Token, error)
	RevokeToken(ctx context.Context, tokenID string) error
}

type oauthRepository struct {
	db *gorm.DB
}

func NewOAuthRepository(db *gorm.DB) OAuthRepository {
	return &oauthRepository{db: db}
}
