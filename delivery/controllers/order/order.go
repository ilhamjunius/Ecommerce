package order

import (
	"ecommerce/delivery/common"
	"ecommerce/entities"
	order "ecommerce/repository/orders"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/invoice"
)

type OrderController struct {
	Repo order.OrderInterface
}

func NewOrderControllers(oi order.OrderInterface) *OrderController {
	return &OrderController{Repo: oi}
}

func (oc OrderController) GetAllOrdersCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		orders, err := oc.Repo.GetAll(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		response := GetOrdersResponseFormat{
			Message: "Successful Operation",
			Data:    orders,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (oc OrderController) GetOrderCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		oid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		order, err := oc.Repo.Get(oid, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		response := GetOrderResponseFormat{
			Message: "Successful Operation",
			Data:    order,
		}
		return c.JSON(http.StatusOK, response)
	}
}

func (oc OrderController) CreateOrderCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		newOrder := entities.Order{
			UserID: uint(id),
			Status: "Open",
		}
		res, err := oc.Repo.Create(newOrder)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}
		return c.JSON(http.StatusOK, CreateOrderResponseFormat{
			Message: "successfull Operation ",
			Data:    res,
		})

	}

}

func (oc OrderController) CancelOrderCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		oid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		res, err := oc.Repo.Cancel(oid, id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, CancelOrderResponseFormat{
			Message: "successfull Operation ",
			Data:    res,
		})
	}
}

func (oc OrderController) PayOrderCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		oid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		order, err := oc.Repo.Get(oid, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		if order.Status != "Open" {
			return c.JSON(http.StatusNotAcceptable, common.NewStatusNotAcceptable())
		}

		// PASTIKAN ORDER ID DAN USER ID PUNYA ORANG TERSEBUT DAN ADA DI DALAM DATABASE DAN DALAM KEADAAN OPEN

		xendit.Opt.SecretKey = "xnd_development_VCICCONHPKS9PAXiekMZBWEyEKPDhRERS3YQZaZ29oZaIfGnSj1HFXErg3kAWcz"

		data := invoice.CreateParams{
			ExternalID: fmt.Sprintf("demoorderID %d", order.ID),
			Amount:     float64(order.Total),
			PayerEmail: fmt.Sprintf("%d", order.UserID),
		}

		resp, err := invoice.Create(&data)
		if err != nil {
			log.Fatal(err)
		}

		res, err := oc.Repo.Pay(resp.ID, resp.InvoiceURL, oid, id)
		if err != nil {
			return c.JSON(http.StatusNotFound, common.NewNotFoundResponse())
		}

		return c.JSON(http.StatusOK, CancelOrderResponseFormat{
			Message: "successfull Operation ",
			Data:    res,
		})
	}
}

func (oc OrderController) CheckOrderCtrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		oid, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, common.NewBadRequestResponse())
		}

		uid := c.Get("user").(*jwt.Token)
		claims := uid.Claims.(jwt.MapClaims)
		id := int(claims["userid"].(float64))

		order, err := oc.Repo.Get(oid, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.NewInternalServerErrorResponse())
		}

		// PASTIKAN ORDER ID DAN USER ID PUNYA ORANG TERSEBUT DAN ADA DI DALAM DATABASE

		xendit.Opt.SecretKey = "xnd_development_VCICCONHPKS9PAXiekMZBWEyEKPDhRERS3YQZaZ29oZaIfGnSj1HFXErg3kAWcz"

		data := invoice.GetParams{
			ID: order.InvoiceID,
		}

		resp, err := invoice.Get(&data)
		if err != nil {
			log.Fatal(err)
		}

		if resp.Status == "PENDING" {
			res, _ := oc.Repo.Get(oid, id)
			response := GetOrderResponseFormat{
				Message: "Successful Operation",
				Data:    res,
			}
			return c.JSON(http.StatusOK, response)
		} else {
			res, _ := oc.Repo.Check(oid, id)
			response := GetOrderResponseFormat{
				Message: "Successful Operation",
				Data:    res,
			}
			return c.JSON(http.StatusOK, response)
		}
	}
}
