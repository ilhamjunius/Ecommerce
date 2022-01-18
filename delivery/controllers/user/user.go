package user

import (
	"crypto/sha256"
	"ecommerce/delivery/common"
	"ecommerce/entities"
	user "ecommerce/repository/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UsersController struct {
	Repo user.UserInterface
}

func NewUsersControllers(ui user.UserInterface) *UsersController {
	return &UsersController{Repo: ui}
}

func (uc UsersController) GetAllUsersCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		user, err := uc.Repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := GetUsersResponseFormat{
			Message: "Successful Operation",
			Data:    user,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (uc UsersController) GetUserCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		user, err := uc.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success",
			"data":    user,
		})
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
			Password:        hash,
			Name:            newUserReq.Name,
			HandphoneNumber: newUserReq.HandphoneNumber,
		}

		_, err := uc.Repo.Create(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

func (uc UsersController) UpdateUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateUserReq := PutUserRequestFormat{}
		if err := c.Bind(&updateUserReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		hash := sha256.Sum256([]byte(updateUserReq.Password))

		updateUser := entities.User{
			Email:           updateUserReq.Email,
			Password:        hash,
			Name:            updateUserReq.Name,
			HandphoneNumber: updateUserReq.HandphoneNumber,
		}

		if _, err := uc.Repo.Update(updateUser, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

func (uc UsersController) DeleteUserCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		deletedUser, _ := uc.Repo.Delete(id)

		if deletedUser.ID != 0 {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}
}
