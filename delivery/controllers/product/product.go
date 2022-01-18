package product

import (
	"ecommerce/delivery/common"
	product "ecommerce/repository/products"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	Repo product.ProductInterface
}

func NewProductControllers(pi product.ProductInterface) *ProductController {
	return &ProductController{Repo: pi}
}

func (pc ProductController) GetAllProductCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		product, err := pc.Repo.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		response := GetAllProductsResponseFormat{
			Message: "Successful Operation",
			Data:    product,
		}
		return c.JSON(http.StatusOK, response)
	}

}
