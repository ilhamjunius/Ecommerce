package auth

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	var user entities.User

	if err := ar.db.Where("Email = ? AND Password=?", email, password).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
