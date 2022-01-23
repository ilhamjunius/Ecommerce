package category

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
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/category")

		categoryController := NewCategoryControllers(mockCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.PostCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Input Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/categoryy")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.PostCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Internal Server Error", response.Message)

	})
	t.Run("Test update Category", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.UpdateCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error update Category", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.UpdateCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Delete Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.DeleteCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Delete Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.DeleteCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Get All Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")

		categoryController := NewCategoryControllers(mockCategoryRepository{})

		categoryController.GetAllCategoryCtrl()(context)
		response := GetCategoriesResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error Get All Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/category")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		categoryController.GetAllCategoryCtrl()(context)
		response := GetCategoriesResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test GetByID Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockCategoryRepository{})

		categoryController.GetCategoryCtrl()(context)

		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error GetByID Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		categoryController.GetCategoryCtrl()(context)

		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
	t.Run("Test Error GetByID Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("asd")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		categoryController.GetCategoryCtrl()(context)

		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Bad Request", response.Message)

	})
}
func TestUserCategory(t *testing.T) {

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

		authControl := auth.NewAuthControllers(mockFalseAuthRepository{})
		authControl.LoginAuthCtrl()(context)

		responses := auth.LoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		// fmt.Println(jwtToken)
		assert.Equal(t, "Successful Operation", responses.Message)
	})
	t.Run("Test Error Input Create", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/category")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.PostCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Authorized", response.Message)

	})
	t.Run("Test Error update Category", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"category_type": "Food",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.UpdateCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Authorized", response.Message)

	})
	t.Run("Test Error Delete Category", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/category:id")
		context.SetParamNames("id")
		context.SetParamValues("0")

		categoryController := NewCategoryControllers(mockFalseCategoryRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(categoryController.DeleteCategoryCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := CreateCategoryResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Authorized", response.Message)

	})

}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "junius123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{Email: "junius@gmail.com", Password: hash[:], Name: "junius", HandphoneNumber: "9172390127390", Role: "admin"}, nil
}

type mockFalseAuthRepository struct{}

func (ma mockFalseAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	passwordHash := "ilham123"
	hash := sha256.Sum256([]byte(passwordHash))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], Name: "ilham", HandphoneNumber: "9721231247390", Role: "user"}, nil
}

type mockCategoryRepository struct{}

func (m mockCategoryRepository) Create(newCategory entities.Category) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, nil
}
func (m mockCategoryRepository) Update(newCategory entities.Category, categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, nil
}

func (m mockCategoryRepository) GetAll() ([]entities.Category, error) {
	return []entities.Category{
		{CategoryType: "Food"},
	}, nil
}
func (m mockCategoryRepository) Get(categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, nil
}
func (m mockCategoryRepository) Delete(categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, nil
}

type mockFalseCategoryRepository struct{}

func (m mockFalseCategoryRepository) Create(newCategory entities.Category) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, errors.New("False Login Object")
}
func (m mockFalseCategoryRepository) Update(newCategory entities.Category, categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, errors.New("False Login Object")
}

func (m mockFalseCategoryRepository) GetAll() ([]entities.Category, error) {
	return []entities.Category{
		{CategoryType: "Food"},
	}, errors.New("False Login Object")
}
func (m mockFalseCategoryRepository) Get(categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, errors.New("False Login Object")
}
func (m mockFalseCategoryRepository) Delete(categoryId int) (entities.Category, error) {
	return entities.Category{CategoryType: "Food"}, errors.New("False Login Object")
}
