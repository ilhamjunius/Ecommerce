package auth

import (
	"crypto/sha256"
	"ecommerce/entities"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) LoginUser(email, password string) (entities.User, error) {
	var user entities.User
	hash := sha256.Sum256([]byte(password))

	ar.db.Where("Email = ? AND Password=?", email, hash).Find(&user)

	return user, nil
}
