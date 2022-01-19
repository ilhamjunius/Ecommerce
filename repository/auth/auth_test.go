package auth

import (
	"crypto/sha256"
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.AutoMigrate(&entities.User{})

	authRepo := NewAuthRepo(db)

	t.Run("Login User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte("andrew123"))
		var mockUser entities.User
		mockUser.Email = "andrew@yahoo.com"
		mockUser.Password = password[:]

		res, err := authRepo.LoginUser(mockUser.Email, mockUser.Password)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Email, res.Email)
	})

	// t.Run("Error Login User into Database", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte(""))
	// 	var mockUser entities.User
	// 	mockUser.Email = ""
	// 	mockUser.Password = password[:]

	// 	_, err := authRepo.LoginUser(mockUser.Email, mockUser.Password)
	// 	assert.NotNil(t, err)
	// })
}
