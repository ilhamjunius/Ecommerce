package category

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	"ecommerce/repository/category"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	Repo category.CategoryInterface
}

func NewCategoryControllers(ci category.CategoryInterface) *CategoryController {
	return &CategoryController{Repo: ci}
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

		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
