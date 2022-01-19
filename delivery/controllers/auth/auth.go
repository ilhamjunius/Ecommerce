package auth

import (
	"crypto/sha256"
	"ecommerce/delivery/common"
	"ecommerce/repository/auth"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	Repo auth.AuthInterface
}

func NewAuthControllers(authrepo auth.AuthInterface) *AuthController {
	return &AuthController{
		Repo: authrepo,
	}
}

func (authcon AuthController) LoginAuthCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequestFormat{}
		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(loginFormat.Password))

		checkedUser, err := authcon.Repo.LoginUser(loginFormat.Email, hash[:])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		token, err := CreateTokenAuth(checkedUser.ID, checkedUser.Role)
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successful Operation",
			"token":   token,
		},
		)
	}
}

func CreateTokenAuth(id uint, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userid"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("RAHASIA"))
}
