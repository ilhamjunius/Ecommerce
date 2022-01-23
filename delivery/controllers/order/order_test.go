package order

import (
	"bytes"
	"crypto/sha256"
	"ecommerce/delivery/controllers/auth"
	"ecommerce/entities"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var jwtToken string

func TestAdminCategory(t *testing.T) {

	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@gmail.com",
			"password": "ilham123",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/login")

		authControl := auth.NewAuthControllers(mockAuthRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := auth.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Input Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string][]int{
			"cartid": []int{1, 2},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/orders")

		orderController := NewOrderControllers(mockOrderRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CreateOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Get All Orders", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders")

		orderController := NewOrderControllers(mockOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.GetAllOrdersCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Get By ID Order", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.GetOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Cancel Order by ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/cancel/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CancelOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Pay Order by ID ", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/pay/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.PayOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Check Order by ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/check/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CheckOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)
	})
	//ERROR TESTING
	t.Run("Test Error Input Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string][]int{
			"cartid": []int{1, 2},
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/orders")

		orderController := NewOrderControllers(mockFalseOrderRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CreateOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Get All Orders", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.GetAllOrdersCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Get Order", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.GetOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Get Order id string", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.GetOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("ErrorTest Cancel Order by ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/cancel/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CancelOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Error Test Cancel Order id string", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/cancel/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CancelOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Error Test Pay Order by ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/pay/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.PayOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Pay Order by ID item waiting for payment", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/pay/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockFalseOrderRepository2{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.PayOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Pay Order id string", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/pay/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.PayOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
	t.Run("Error Test Check Order by ID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/check/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CheckOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetOrderResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Error Test Check Order id string", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/orders/pay/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		orderController := NewOrderControllers(mockFalseOrderRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(orderController.CheckOrderCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetOrdersResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)
	})
}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "junius123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{Email: "junius@gmail.com", Password: hash[:], Name: "junius", HandphoneNumber: "9172390127390", Role: "admin"}, nil
}

type mockOrderRepository struct{}

func (m mockOrderRepository) GetAll(userid int) ([]entities.Order, error) {
	return []entities.Order{
		{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"},
	}, nil
}
func (m mockOrderRepository) Get(orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, nil
}

func (m mockOrderRepository) Create(newOrder entities.Order, arr []int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, nil
}
func (m mockOrderRepository) Cancel(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "Cancel", PaymentLink: "Cancel", UserID: 1, Total: 5000, Status: "Cancel"}, nil
}
func (m mockOrderRepository) Pay(invoiceId, paymentLink string, orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, nil
}

func (m mockOrderRepository) Check(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, nil
}

type mockFalseOrderRepository struct{}

func (m mockFalseOrderRepository) GetAll(userid int) ([]entities.Order, error) {
	return []entities.Order{
		{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"},
	}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository) Get(orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, errors.New("False Login Object")
}

func (m mockFalseOrderRepository) Create(newOrder entities.Order, arr []int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository) Cancel(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "Cancel", PaymentLink: "Cancel", UserID: 1, Total: 5000, Status: "Cancel"}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository) Pay(invoiceId, paymentLink string, orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, errors.New("False Login Object")
}

func (m mockFalseOrderRepository) Check(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, errors.New("False Login Object")
}

type mockFalseOrderRepository2 struct{}

func (m mockFalseOrderRepository2) GetAll(userid int) ([]entities.Order, error) {
	return []entities.Order{
		{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"},
	}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository2) Get(orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, errors.New("False Login Object")
}

func (m mockFalseOrderRepository2) Create(newOrder entities.Order, arr []int) (entities.Order, error) {
	return entities.Order{InvoiceID: "", PaymentLink: "", UserID: 1, Total: 5000, Status: "Open"}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository2) Cancel(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "Cancel", PaymentLink: "Cancel", UserID: 1, Total: 5000, Status: "Cancel"}, errors.New("False Login Object")
}
func (m mockFalseOrderRepository2) Pay(invoiceId, paymentLink string, orderId, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, nil
}

func (m mockFalseOrderRepository2) Check(orderId int, userId int) (entities.Order, error) {
	return entities.Order{InvoiceID: "demo1", PaymentLink: "xendit.com", UserID: 1, Total: 5000, Status: "Waiting for Payment"}, errors.New("False Login Object")
}
