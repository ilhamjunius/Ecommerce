package user

import (
	"crypto/sha256"
	"ecommerce/delivery/common"
	"ecommerce/entities"
	user "ecommerce/repository/users"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo user.UserInterface
}

func NewUsersControllers(ui user.UserInterface) *UsersController {
	return &UsersController{Repo: ui}
}

func (uc UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		user, err := uc.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := GetUserResponseFormat{
			Message: "Successful Operation",
			Data:    user,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (uc UsersController) RegisterUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		newUserReq := RegisterUserRequestFormat{}
		if err := c.Bind(&newUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(newUserReq.Password))

		newUser := entities.User{
			Email:           newUserReq.Email,
			Password:        hash[:],
			Name:            newUserReq.Name,
			HandphoneNumber: newUserReq.HandphoneNumber,
			Role:            newUserReq.Role,
		}

		_, err := uc.Repo.Create(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := RegisterUserResponseFormat{
			Message: "Successful Operation",
			Data:    newUser,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (uc UsersController) UpdateUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		updateUserReq := PutUserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateUserReq.Password))

		updateUser := entities.User{
			Email:           updateUserReq.Email,
			Password:        hash[:],
			Name:            updateUserReq.Name,
			HandphoneNumber: updateUserReq.HandphoneNumber,
			Role:            updateUserReq.Role,
		}

		if _, err := uc.Repo.Update(updateUser, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := PutUserResponseFormat{
			Message: "Successful Operation",
			Data:    updateUser,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (uc UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))
		deletedUser, err := uc.Repo.Delete(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := DeleteUserResponseFormat{
			Message: "Successful Operation",
			Data:    deletedUser,
		}

		return c.JSON(http.StatusOK, response)
	}
}
