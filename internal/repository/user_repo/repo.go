package user_repo

import (
	"context"
	"github.com/riyanathariq/taskify-api/internal/models"
)

func (r *userRepo) Create(ctx context.Context, payload *models.User) error {
	return r.db.WithContext(ctx).Create(payload).Error
}

func (r *userRepo) DetailByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) DetailByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) List(ctx context.Context, filter *models.User) ([]*models.User, error) {
	var users []*models.User
	query := r.db.WithContext(ctx).Model(&models.User{})

	if filter.FilterData().Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.FilterData().Name+"%")
	}
	if filter.FilterData().Username != "" {
		query = query.Where("username ILIKE ?", "%"+filter.FilterData().Username+"%")
	}

	if filter.FilterData().Limit > 0 {
		query = query.Limit(filter.FilterData().Limit)
	}
	if filter.FilterData().Offset > 0 {
		query = query.Offset(filter.FilterData().Offset)
	}

	err := query.Find(&users).Error
	return users, err
}

func (r *userRepo) Update(ctx context.Context, id string, update *models.User) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Updates(update).Error
}

func (r *userRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
}
