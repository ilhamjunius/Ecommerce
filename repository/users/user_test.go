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

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

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
	t.Run("Insert User into Database", func(t *testing.T) {
		password := sha256.Sum256([]byte("andrew123"))
		var mockUser entities.User
		mockUser.Email = "andrew@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "andrew"
		mockUser.HandphoneNumber = "0123456789"
		mockUser.Role = "admin"
		db.Migrator().DropTable(&entities.User{})
		_, err := userRepo.Create(mockUser)
		assert.Error(t, err)
		// assert.Equal(t, mockUser.Name, res.Name)
		// assert.Equal(t, 1, int(res.ID))

	})
	t.Run("Select Users from Database", func(t *testing.T) {
		db.AutoMigrate(&entities.User{})
		res, err := userRepo.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})
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
	t.Run("Select User from Database", func(t *testing.T) {
		_, err := userRepo.Get(1000)
		assert.Error(t, err)

	})

	t.Run("Error Update User ", func(t *testing.T) {
		password := sha256.Sum256([]byte("ilham123"))
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "ilham"
		mockUser.HandphoneNumber = "987654321"
		mockUser.Role = "pembeli"

		_, err := userRepo.Update(mockUser, 100)
		assert.Error(t, err)

	})
	t.Run("Update User ", func(t *testing.T) {
		password := sha256.Sum256([]byte("ilham123"))
		var mockUser entities.User
		mockUser.Email = "ilham@yahoo.com"
		mockUser.Password = password[:]
		mockUser.Name = "ilham"
		mockUser.HandphoneNumber = "987654321"
		mockUser.Role = "pembeli"

		res, err := userRepo.Update(mockUser, 1)
		assert.Nil(t, err)
		assert.Equal(t, mockUser.Name, res.Name)
	})

	t.Run("Error Delete User", func(t *testing.T) {
		_, err := userRepo.Delete(100)
		assert.Error(t, err)
	})
	t.Run("Delete User", func(t *testing.T) {
		res, err := userRepo.Delete(1)
		assert.Nil(t, err)
		assert.Equal(t, res, res)
	})

	// t.Run("Insert User into Database", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("andrew123"))
	// 	var mockUser entities.User
	// 	mockUser.Email = "andrew@yahoo.com"
	// 	mockUser.Password = password[:]
	// 	mockUser.Name = "andrew"
	// 	mockUser.HandphoneNumber = "0123456789"
	// 	mockUser.Role = "admin"

	// 	res, err := userRepo.Create(mockUser)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, mockUser.Name, res.Name)
	// 	assert.Equal(t, 2, int(res.ID))
	// })

	// t.Run("Update User ", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("ilham123"))
	// 	var mockUser entities.User
	// 	mockUser.Email = "ilham@yahoo.com"
	// 	mockUser.Password = password[:]
	// 	mockUser.Name = "ilham"
	// 	mockUser.HandphoneNumber = "987654321"
	// 	mockUser.Role = "pembeli"

	// 	res, err := userRepo.Create(mockUser)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, mockUser.Name, res.Name)
	// })

	// t.Run("Error Data Update User ", func(t *testing.T) {
	// 	var mockUser entities.User
	// 	mockUser.Email = "andrew@yahoo.com"

	// 	_, err := userRepo.Update(mockUser, 3)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Error ID Update User ", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("ilham123"))
	// 	var mockUpdateUser entities.User
	// 	mockUpdateUser.Email = "ilham@yahoo.com"
	// 	mockUpdateUser.Password = password[:]
	// 	mockUpdateUser.Name = "ilham"
	// 	mockUpdateUser.HandphoneNumber = "987654321"
	// 	mockUpdateUser.Role = "pembeli"

	// 	_, err := userRepo.Update(mockUpdateUser, 10)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Error Delete User", func(t *testing.T) {
	// 	_, err := userRepo.Delete(10)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Insert User into Database", func(t *testing.T) {
	// 	password := sha256.Sum256([]byte("andrew123"))
	// 	var mockInsertUser entities.User
	// 	mockInsertUser.Email = "andrew@yahoo.com"
	// 	mockInsertUser.Password = password[:]
	// 	mockInsertUser.Name = "andrew"
	// 	mockInsertUser.HandphoneNumber = "0123456789"
	// 	mockInsertUser.Role = "admin"

	// 	_, err := userRepo.Create(mockInsertUser)
	// 	assert.Nil(t, err)
	// })
}
