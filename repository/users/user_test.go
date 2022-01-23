package user

import (
	"crypto/sha256"
	"ecommerce/configs"
	"ecommerce/entities"
	"ecommerce/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersRepo(t *testing.T) {
	config := configs.GetConfig()
	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Order{})
	db.AutoMigrate(&entities.Category{})
	db.AutoMigrate(&entities.Product{})
	db.AutoMigrate(&entities.ShoppingCart{})

	userRepo := NewUserRepo(db)

	t.Run("Insert User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte("andrew123"))
		var mockUser entities.User
		mockUser.Email = "andrew@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "andrew"
		mockUser.HandphoneNumber = "0123456789"
		mockUser.Role = "admin"

		res, err := userRepo.Create(mockUser)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, 1, int(res.ID))

	})

	t.Run("Select User from Database", func(t *testing.T) {
		res, err := userRepo.Get(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	t.Run("Update User ", func(t *testing.T) {
		password := sha256.Sum256([]byte("ilham123"))
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "ilham"
		mockUser.HandphoneNumber = "987654321"
		mockUser.Role = "admin"

		res, err := userRepo.Update(mockUser, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("Delete User", func(t *testing.T) {
		res, err := userRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	//DropTable
	db.Migrator().DropTable(&entities.ShoppingCart{})
	db.Migrator().DropTable(&entities.Order{})
	db.Migrator().DropTable(&entities.Product{})
	db.Migrator().DropTable(&entities.Category{})
	db.Migrator().DropTable(&entities.User{})

	t.Run("Error Insert User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte("andrew123"))
		var mockUser entities.User
		mockUser.Email = "andrew@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "andrew"
		mockUser.HandphoneNumber = "0123456789"
		mockUser.Role = "admin"

		_, err := userRepo.Create(mockUser)
		assert.NotNil(t, err)
	})

	t.Run("Error Select User from Database", func(t *testing.T) {
		_, err := userRepo.Get(1)
		assert.NotNil(t, err)
	})

	t.Run("Error Update User ", func(t *testing.T) {
		password := sha256.Sum256([]byte("ilham123"))
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "ilham"
		mockUser.HandphoneNumber = "987654321"
		mockUser.Role = "admin"

		_, err := userRepo.Update(mockUser, 1)
		assert.NotNil(t, err)
	})

	t.Run("Delete User", func(t *testing.T) {
		_, err := userRepo.Delete(1)
		assert.NotNil(t, err)
	})
}
