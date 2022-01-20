package user

import (
	"ecommerce/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Create(newUser entities.User) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.Find(&user, "email=?", user.Email).Error; err != nil {
		return newUser, err
	}
	ur.db.Save(&newUser)

	return newUser, nil
}

func (ur *UserRepository) Update(newUser entities.User, userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.First(&user, "id=?", userId).Error; err != nil {
		return newUser, err
	}

	user.Email = newUser.Email
	user.Password = newUser.Password
	user.Name = newUser.Name
	user.HandphoneNumber = newUser.HandphoneNumber
	user.Role = newUser.Role

	ur.db.Save(&user)

	return newUser, nil
}

func (ur *UserRepository) Delete(userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.First(&user, "id=?", userId).Error; err != nil {
		return user, err
	}
	ur.db.Delete(&user)

	return user, nil
}
