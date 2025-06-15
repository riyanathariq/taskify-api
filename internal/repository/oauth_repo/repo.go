package oauth_repo

import (
	"context"
	"github.com/riyanathariq/taskify-api/internal/models"
)

func (r *oauthRepository) SaveToken(ctx context.Context, token *models.Token) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *oauthRepository) FindTokenByAccessToken(ctx context.Context, accessToken string) (*models.Token, error) {
	var token models.Token
	err := r.db.WithContext(ctx).
		Where("access_token = ? AND revoked = false", accessToken).
		First(&token).Error
	return &token, err
}

func (r *oauthRepository) FindTokenByRefreshToken(ctx context.Context, refreshToken string) (*models.Token, error) {
	var token models.Token
	err := r.db.WithContext(ctx).
		Where("refresh_token = ? AND revoked = false", refreshToken).
		First(&token).Error
	return &token, err
}

func (r *oauthRepository) RevokeToken(ctx context.Context, tokenID string) error {
	return r.db.WithContext(ctx).
		Model(&models.Token{}).
		Where("id = ?", tokenID).
		Update("revoked", true).Error
}
