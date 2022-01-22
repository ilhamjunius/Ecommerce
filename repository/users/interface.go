package user

import "ecommerce/entities"

type UserInterface interface {
	Get(userId int) (entities.User, error)
	Create(newUser entities.User) (entities.User, error)
	Update(newUser entities.User, userId int) (entities.User, error)
	Delete(userId int) (entities.User, error)
}
