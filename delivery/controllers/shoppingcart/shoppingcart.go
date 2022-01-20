package shoppingcart

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	"ecommerce/repository/shoppingcart"
	"net/http"

	"github.com/golang-jwt/jwt"
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
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		newShoppingCartreq := ShoppingCartRequestFormat{}
		if err := c.Bind(&newShoppingCartreq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}
		newShoppingCart := entities.ShoppingCart{
			OrderId:   newShoppingCartreq.OrderId,
			UserID:    uint(id),
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

		// id, err := strconv.Atoi(c.Param("id"))
		// if err != nil {
		// 	return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		// }
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		updateShoppingCartReq := ShoppingCartRequestFormat{}
		if err := c.Bind(&updateShoppingCartReq); err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		updateShopingCart := entities.ShoppingCart{
			Qty:     updateShoppingCartReq.Qty,
			OrderId: updateShoppingCartReq.OrderId,
		}

		if _, err := sc.Repo.Update(updateShopingCart, id); err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
		return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
	}
}

func (sc ShoppingCartController) DeleteShoppingCartCtrl() echo.HandlerFunc {

	return func(c echo.Context) error {

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		deletedShoppingCart, _ := sc.Repo.Delete(id)

		if deletedShoppingCart.ID != 0 {
			return c.JSON(http.StatusOK, common.NewSuccessOperationResponse())
		} else {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}
	}
}
