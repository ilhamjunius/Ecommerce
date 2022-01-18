package auth

import "ecommerce/entities"

type AuthInterface interface {
	LoginUser(email string, password [32]byte) (entities.User, error)
}
