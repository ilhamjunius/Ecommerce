package auth

import "ecommerce/entities"

type AuthInterface interface {
	LoginUser(email, password string) (entities.User, error)
}
