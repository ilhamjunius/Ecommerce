package shoppingcart

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

func TestShoppingCart(t *testing.T) {
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "junius@gmail.com",
			"password": "junius123",
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
		// fmt.Println(jwtToken)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Error GetByUserID ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart")

		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.GetShppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ManyShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test GetByUserID ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart")

		shoppingCartController := NewShoppingCartControllers(mockShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.GetShppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ManyShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successfull Operation", response.Message)

	})
	t.Run("Test Error Create ShoppingCarts", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"product_id": 1,
			"quantity":   100,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart")

		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.CreateShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test Create ShoppingCarts", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"product_id": 1,
			"quantity":   100,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart")

		shoppingCartController := NewShoppingCartControllers(mockShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.CreateShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successfull Operation", response.Message)

	})
	t.Run("Test Error Update ShoppingCarts", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"user_id":  1,
			"order_id": 1,
			"quantity": 101,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.UpdateShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		// fmt.Println(jwtToken)
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Update CartId ShoppingCarts", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"order_id": 1,
			"quantity": 101,
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")
		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.UpdateShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Update ShoppingCarts", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]int{
			"user_id":  1,
			"quantity": 101,
			"order_id": 1,
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		shoppingCartController := NewShoppingCartControllers(mockShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.UpdateShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Delete ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("5")
		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.DeleteShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Delete CartId ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		shoppingCartController := NewShoppingCartControllers(mockShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.DeleteShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Delete ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("2")

		shoppingCartController := NewShoppingCartControllers(mockShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.DeleteShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println("Response:", response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Delete ShoppingCarts", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/shoppingcart/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		shoppingCartController := NewShoppingCartControllers(mockFalseShoppingCartRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(shoppingCartController.DeleteShoppingCartCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ShoppingCartResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println("Response:", response)
		assert.Equal(t, "Not Found", response.Message)

	})
}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "junius123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{ID: 1, Email: "junius@gmail.com", Password: hash[:], Name: "junius", HandphoneNumber: "9172390127390", Role: "admin"}, nil
}

type mockShoppingCartRepository struct{}

func (m mockShoppingCartRepository) Get(userId int) ([]entities.ShoppingCart, error) {
	return []entities.ShoppingCart{
		{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000},
	}, nil
}
func (m mockShoppingCartRepository) GetById(id, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, nil
}
func (m mockShoppingCartRepository) Update(updateCart entities.ShoppingCart, cartId, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, OrderId: updateCart.OrderId, Qty: updateCart.Qty, Subtotal: 10000}, nil
}

func (m mockShoppingCartRepository) Create(newCart entities.ShoppingCart) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, nil
}
func (m mockShoppingCartRepository) Delete(cartId, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, nil
}

type mockFalseShoppingCartRepository struct{}

func (m mockFalseShoppingCartRepository) Get(userId int) ([]entities.ShoppingCart, error) {
	return []entities.ShoppingCart{
		{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000},
	}, errors.New("False Login Object")
}
func (m mockFalseShoppingCartRepository) GetById(id, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, errors.New("False Login Object")
}
func (m mockFalseShoppingCartRepository) Update(updateCart entities.ShoppingCart, cartId, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, UserID: 1, OrderId: 1, Qty: 100}, errors.New("False Login Object")
}

func (m mockFalseShoppingCartRepository) Create(newCart entities.ShoppingCart) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, errors.New("False Login Object")
}
func (m mockFalseShoppingCartRepository) Delete(cartId, userId int) (entities.ShoppingCart, error) {
	return entities.ShoppingCart{ID: 1, ProductID: 1, UserID: 1, Qty: 100, Subtotal: 10000}, errors.New("False Login Object")
}
