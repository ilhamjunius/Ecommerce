package product

import (
	"bytes"
	"crypto/sha256"
	"ecommerce/delivery/controllers/auth"
	"ecommerce/entities"
	product "ecommerce/repository/products"
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

func TestProduct(t *testing.T) {
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
	t.Run("Test Error GetProductByID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		productController := NewProductControllers(mockProductRepository{})

		productController.GetProductCtrl()(context)
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Error GetProductByID", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		productController := NewProductControllers(mockFalseProductRepository{})

		productController.GetProductCtrl()(context)
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test GetProductByID ", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")

		productController := NewProductControllers(mockProductRepository{})

		productController.GetProductCtrl()(context)
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Product Pagination ", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("name")
		context.SetParamValues("mie")
		context.SetParamNames("category")
		context.SetParamValues("mie")
		context.SetParamNames("limit")
		context.SetParamValues("1")
		context.SetParamNames("page")
		context.SetParamValues("0")
		context.SetParamNames("sort")
		context.SetParamValues("id asc")

		productController := NewProductControllers(mockProductRepository{})

		productController.Pagination()(context)
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Product Pagination ", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/products")
		context.SetParamNames("name")
		context.SetParamValues("mie")
		context.SetParamNames("category")
		context.SetParamValues("mie")
		context.SetParamNames("limit")
		context.SetParamValues("1")
		context.SetParamNames("page")
		context.SetParamValues("0")
		context.SetParamNames("sort")
		context.SetParamValues("id asc")

		productController := NewProductControllers(mockFalseProductRepository{})

		productController.Pagination()(context)
		response := PaginationResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Error Create Product", func(t *testing.T) {
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
		context.SetPath("/products")

		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.CreateProductControllers())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test Create Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.CreateProductControllers())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successfull Operation", response.Message)

	})
	t.Run("Test Error Update Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.UpdateProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Id Update Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.UpdateProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Bad Request", response.Message)

	})

	t.Run("Test Update Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.UpdateProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Delete Product", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("100")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.DeleteProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Id Delete Product", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.DeleteProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Delete Product", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.DeleteProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Update Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.UpdateProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Successful Operation", response.Message)

	})

}
func TestUserProduct(t *testing.T) {
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

		authControl := auth.NewAuthControllers(mockFalseAuthRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := auth.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		// fmt.Println(jwtToken)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Create Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products")

		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.CreateProductControllers())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Not Authorized", response.Message)

	})
	t.Run("Test Update Product", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"product_name": "mie goreng",
			"price":        3000,
			"stock":        100,
			"category_id":  1,
			"description":  "mie goreng spesial",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.UpdateProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Not Authorized", response.Message)

	})
	t.Run("Test Delete Product", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/products/:id")
		context.SetParamNames("id")
		context.SetParamValues("1")
		productController := NewProductControllers(mockFalseProductRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(productController.DeleteProductCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := ProductResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// fmt.Println(response)
		assert.Equal(t, "Not Authorized", response.Message)

	})
}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "junius123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{ID: 1, Email: "junius@gmail.com", Password: hash[:], Name: "junius", HandphoneNumber: "9172390127390", Role: "admin"}, nil
}

type mockFalseAuthRepository struct{}

func (ma mockFalseAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "ilham123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{ID: 1, Email: "ilham@gmail.com", Password: hash[:], Name: "ilham", HandphoneNumber: "917232347390", Role: "user"}, nil
}

type mockProductRepository struct{}

func (m mockProductRepository) Pagination(name, category string, pagination entities.Pagination) (product.ResultPagination, int) {
	return product.ResultPagination{
		Result: entities.Pagination{
			Limit:     pagination.Limit,
			Page:      pagination.Page,
			Sort:      pagination.Sort,
			TotalRows: pagination.TotalRows,
			FromRow:   pagination.FromRow,
			ToRow:     pagination.ToRow,
			Rows:      pagination.Rows,
		},
		Error: nil,
	}, 0
}
func (m mockProductRepository) Get(productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, nil
}

func (m mockProductRepository) Create(newProduct entities.Product) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, nil
}
func (m mockProductRepository) Delete(productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, nil
}
func (m mockProductRepository) Update(newProduct entities.Product, productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, nil
}

type mockFalseProductRepository struct{}

func (m mockFalseProductRepository) Pagination(name, category string, pagination entities.Pagination) (product.ResultPagination, int) {
	return product.ResultPagination{
		Result: entities.Pagination{
			Limit:     pagination.Limit,
			Page:      pagination.Page,
			Sort:      pagination.Sort,
			TotalRows: pagination.TotalRows,
			FromRow:   pagination.FromRow,
			ToRow:     pagination.ToRow,
			Rows:      pagination.Rows,
		},
		Error: errors.New("False Login Object"),
	}, 0
}
func (m mockFalseProductRepository) Get(productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, errors.New("False Login Object")
}

func (m mockFalseProductRepository) Create(newProduct entities.Product) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, errors.New("False Login Object")
}
func (m mockFalseProductRepository) Delete(productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, errors.New("False Login Object")
}
func (m mockFalseProductRepository) Update(newProduct entities.Product, productId int) (entities.Product, error) {
	return entities.Product{ID: 1, Name: "mie goreng", Price: 3000, Stock: 100, Description: "mie goreng spesial", CategoryID: 1}, errors.New("False Login Object")
}
