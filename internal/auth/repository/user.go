package repository

import (
	"boilerplate-clean-arch/models"
	"context"

	"github.com/google/uuid"
)

func (a *authRepo) SignUp(ctx context.Context, user *models.User) (*models.User, error) {
	if result := a.db.Create(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (a *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	if result := a.db.Where("email = ?", user.Email).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (a *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	var user models.User
	if result := a.db.Where("id = ?", userID).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
