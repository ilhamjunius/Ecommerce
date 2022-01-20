package product

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	product "ecommerce/repository/products"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
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

func (pc ProductController) GetProductCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		product, err := pc.Repo.Get(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		response := ProductResponseFormat{
			Message: "Successful Operation",
			Data:    product,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (pc ProductController) CreateProductControllers() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		newProductreq := ProductRequestFormat{}

		if err := c.Bind(&newProductreq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newProduct := entities.Product{
			Name:        newProductreq.Name,
			Price:       newProductreq.Price,
			Stock:       newProductreq.Stock,
			CategoryID:  newProductreq.CategoryID,
			Description: newProductreq.Description,
		}
		res, err := pc.Repo.Create(newProduct)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		return c.JSON(http.StatusOK, ProductResponseFormat{
			Message: "success create new product",
			Data:    res,
		})
	}
}

func (pc ProductController) UpdateProductCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateProductReq := ProductRequestFormat{}
		if err := c.Bind(&updateProductReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateProduct := entities.Product{
			Name:        updateProductReq.Name,
			Price:       updateProductReq.Price,
			Stock:       updateProductReq.Stock,
			CategoryID:  updateProductReq.CategoryID,
			Description: updateProductReq.Description,
		}

		if _, err := pc.Repo.Update(updateProduct, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}
func (pc ProductController) DeleteProductCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		Role := claims["role"]

		if Role != "admin" {
			return c.JSON(http.StatusBadRequest, common.NewStatusNotAuthorized())
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		deletedProduct, _ := pc.Repo.Delete(id)

		if deletedProduct.ID != 0 {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}
}
