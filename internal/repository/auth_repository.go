package repository

import (
	"github.com/ahmadalaik/be-invitation/internal/models"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Create(user models.User) error
	FindByEmail(email string) (*models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) Create(user models.User) error {
	return r.db.Create(&user).Error
}

func (r *authRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
