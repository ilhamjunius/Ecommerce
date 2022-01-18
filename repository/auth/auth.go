package auth

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) LoginUser(email string, password [32]byte) (entities.User, error) {
	var user entities.User

	if err := ar.db.Where("Email = ? AND Password=?", email, password).Find(&user).Error; err != nil {
		log.Warn("Found database error", err)
		return user, err
	}

	return user, nil
}
