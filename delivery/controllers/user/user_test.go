package user

import (
	"bytes"
	"crypto/sha256"
	"ecommerce/delivery/controllers/auth"
	"ecommerce/entities"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/magiconair/properties/assert"
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
	assert.Equal(t, jwtToken, jwtToken)

	// t.Run("Test Get User", func(t *testing.T) {
	// 	e := echo.New()
	// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	res := httptest.NewRecorder()

	// 	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
	// 	context := e.NewContext(req, res)
	// 	context.SetPath("/users")

	// 	userController := NewUsersControllers(mockUserRepository{})
	// 	if err := userController.GetUserCtrl()(context); err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}

	// 	response := GetUserResponseFormat{}

	// 	json.Unmarshal([]byte(res.Body.Bytes()), &response)
	// 	assert.Equal(t, "Ilham", response.Data.Name)
	// })
}

type mockAuthRepository struct{}

func (ma mockAuthRepository) LoginUser(email string, password []byte) (entities.User, error) {
	hash := sha256.Sum256([]byte("ilham123"))
	return entities.User{Email: "ilham@gmail.com", Password: hash[:], HandphoneNumber: "0123456789", Role: "admin"}, nil
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
