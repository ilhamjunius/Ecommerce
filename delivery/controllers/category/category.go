package category

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	"ecommerce/repository/category"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	Repo category.CategoryInterface
}

func NewCategoryControllers(ci category.CategoryInterface) *CategoryController {
	return &CategoryController{Repo: ci}
}

func (ci CategoryController) GetAllCategoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		categories, err := ci.Repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		response := GetCategoriesResponseFormat{
			Message: "Successful Operation",
			Data:    categories,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (ci CategoryController) GetCategoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		category, err := ci.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		response := GetCategoryResponseFormat{
			Message: "Successful Operation",
			Data:    category,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (ci CategoryController) PostCategoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		newCategoryReq := CreateCategoryRequestFormat{}
		if err := c.Bind(&newCategoryReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newCategory := entities.Category{
			ID:           newCategoryReq.ID,
			CategoryType: newCategoryReq.CategoryType,
		}

		_, err := ci.Repo.Create(newCategory)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := CreateCategoryResponseFormat{
			Message: "Successful Operation",
			Data:    newCategory,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (ci CategoryController) UpdateCategoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]
		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		newCategoryReq := PutCategoryRequestFormat{}
		if err := c.Bind(&newCategoryReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		newCategory := entities.Category{
			ID:           newCategoryReq.ID,
			CategoryType: newCategoryReq.CategoryType,
		}

		if _, err := ci.Repo.Update(newCategory, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := PutCategoryResponseFormat{
			Message: "Successful Operation",
			Data:    newCategory,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func (ci CategoryController) DeleteCategoryCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]
		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		deletedCategory, err := ci.Repo.Delete(id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		response := DeleteCategoryResponseFormat{
			Message: "Successful Operation",
			Data:    deletedCategory,
		}
		return c.JSON(http.StatusOK, response)
	}
}
