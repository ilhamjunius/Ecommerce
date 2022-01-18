package user

import (
	"ecommerce/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetAll() ([]entities.User, error) {
	users := []entities.User{}
	if err := ur.db.Find(&users).Error; err != nil {
		log.Warn("Found database error", err)
		return nil, err
	}

	return users, nil
}

func (ur *UserRepository) Get(userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.Find(&user, userId).Error; err != nil {
		log.Warn("Found database error", err)
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) Create(newUser entities.User) (entities.User, error) {
	if err := ur.db.Save(&newUser).Error; err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (ur *UserRepository) Update(newUser entities.User, userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.Find(&user, "id=?", userId).Error; err != nil {
		return newUser, err
	}

	user = newUser

	if err := ur.db.Save(&user).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (ur *UserRepository) Delete(userId int) (entities.User, error) {
	user := entities.User{}
	if err := ur.db.Find(&user, "id=?", userId).Error; err != nil {
		return user, err
	}
	if err := ur.db.Delete(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
