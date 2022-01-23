package user

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

func TestUser(t *testing.T) {
	t.Run("Test Register", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/userss")

		userController := NewUsersControllers(mockFalseUserRepository{})
		userController.RegisterUserCtrl()(context)

		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Internal Server Error", response.Message)
	})
	t.Run("Test Register", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		userController.RegisterUserCtrl()(context)

		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
		assert.Equal(t, "ilham", response.Data.Name)
		assert.Equal(t, "ilham@yahoo.com", response.Data.Email)
	})

	// LOGIN
	t.Run("Test Error Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "ilham@gmail.com",
			"password": "",
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
		assert.Equal(t, "Internal Server Error", responses.Message)
	})
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

	t.Run("Test GetByID User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(userController.GetUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Successful Operation", response.Message)

	})
	t.Run("Test Error GetByID User", func(t *testing.T) {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})

		if err := middleware.JWT([]byte("RAHASIA"))(userController.GetUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := GetUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Not Found", response.Message)

	})
	t.Run("Test Error Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]string{
			"name":     "ilham",
			"email":    "ilham@yahoo.com",
			"password": "ilham123",
			"role":     "admin",
		})

		req := httptest.NewRequest(http.MethodPut, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.UpdateUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
	t.Run("Test Error Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockFalseUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Not Found", response.Message)
	})
	t.Run("Test Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := NewUsersControllers(mockUserRepository{})
		if err := middleware.JWT([]byte("RAHASIA"))(userController.DeleteUserCtrl())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := RegisterUserResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, "Successful Operation", response.Message)
	})
}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], HandphoneNumber: "0123456789", Role: "admin"}, nil
}

type mockFalseAuthRepository struct{}

func (ma mockFalseAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], HandphoneNumber: "0123456789", Role: "user"}, errors.New("False Login Object")
}

type mockUserRepository struct{}

func (m mockUserRepository) Get(userid int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], HandphoneNumber: "0123456789", Role: "admin"}, nil
}

func (m mockUserRepository) Create(user entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Name: user.Name, Email: user.Email, Password: hash[:], Role: user.Role}, nil
}

func (m mockUserRepository) Update(user entities.User, id int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Name: user.Name, Email: user.Email, Password: hash[:]}, nil
}

func (m mockUserRepository) Delete(id int) (entities.User, error) {
	return entities.User{Name: "Ilham", Email: "ilham@gmail.com"}, nil
}

type mockFalseUserRepository struct{}

func (m mockFalseUserRepository) Get(userid int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], HandphoneNumber: "0123456789", Role: "admin"}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Create(user entities.User) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Name: user.Name, Email: user.Email, Password: hash[:], Role: user.Role}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Update(user entities.User, id int) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Name: user.Name, Email: user.Email, Password: hash[:]}, errors.New("False Login Object")
}

func (m mockFalseUserRepository) Delete(id int) (entities.User, error) {
	return entities.User{Name: "Ilham", Email: "ilham@gmail.com"}, errors.New("False Login Object")
}
