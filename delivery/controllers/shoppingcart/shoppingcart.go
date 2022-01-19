package shoppingcart

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	"ecommerce/repository/shoppingcart"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ShoppingCartController struct {
	Repo shoppingcart.ShoppingCartInterface
}

func NewShoppingCartControllers(si shoppingcart.ShoppingCartInterface) *ShoppingCartController {
	return &ShoppingCartController{Repo: si}
}

func (sc ShoppingCartController) CreateShoppingCartCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		newShoppingCartreq := ShoppingCartRequestFormat{}
		if err := c.Bind(&newShoppingCartreq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newShoppingCart := entities.ShoppingCart{
			UserID:    newShoppingCartreq.UserId,
			ProductID: newShoppingCartreq.ProductId,
			Qty:       newShoppingCartreq.Qty,
		}
		res, err := sc.Repo.Create(newShoppingCart)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		return c.JSON(http.StatusOK, ProductResponseFormat{
			Message: "successfull Operation ",
			Data:    res,
		})

	}

}
func (sc ShoppingCartController) UpdateShoppingCartCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateShoppingCartReq := ShoppingCartRequestFormat{}
		if err := c.Bind(&updateShoppingCartReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateShopingCart := entities.ShoppingCart{
			Qty: updateShoppingCartReq.Qty,
		}

		if _, err := sc.Repo.Update(updateShopingCart.Qty, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

func (sc ShoppingCartController) DeleteShoppingCartCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		deletedShoppingCart, _ := sc.Repo.Delete(id)

		if deletedShoppingCart.ID != 0 {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}
}
