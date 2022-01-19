package auth

import "ecommerce/entities"

type AuthInterface interface {
	LoginUser(email string, password []byte) (entities.User, error)
}
