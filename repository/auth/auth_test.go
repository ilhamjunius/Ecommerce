package auth

import (
	"crypto/sha256"
	"ecommerce/configs"
	"ecommerce/entities"
	user "ecommerce/repository/users"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	userRepo := user.NewUserRepo(db)
	authRepo := NewAuthRepo(db)

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.ShoppingCart{})

	var mockUser entities.User
	hash := sha256.Sum256([]byte("ilham123"))
	mockUser.Email = "ilham@gmail.com"
	mockUser.Password = hash[:]
	mockUser.HandphoneNumber = "0123456789"
	mockUser.Role = "admin"
	userRepo.Create(mockUser)

	t.Run("Login User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte("ilham123"))
		var mockUser entities.User
		mockUser.Email = "ilham@gmail.com"
		mockUser.Password = password[:]

		res, err := authRepo.LoginUser(mockUser.Email, mockUser.Password)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Email, res.Email)
	})

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Error Login User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte(""))
		var mockUser entities.User
		mockUser.Email = ""
		mockUser.Password = password[:]

		_, err := authRepo.LoginUser(mockUser.Email, mockUser.Password)
		assert.NotNil(t, err)
	})
}
