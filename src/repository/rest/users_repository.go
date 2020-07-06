package rest

import (
	"encoding/json"
	"github.com/codewitch24/BookstoreOAuthAPI/src/domain/users"
	"github.com/codewitch24/BookstoreOAuthAPI/src/utils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "http://127.0.0.1:8080",
		Timeout: 100 * time.Millisecond,
	}
)

type UserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestError)
}

type usersRepository struct{}

func NewRepository() UserRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestError) {
	requestBody := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}
	response := usersRestClient.Post("/users/login", requestBody)
	if response == nil || response.Response == nil {
		return nil, errors.NewInternalServerError("Invalid rest client response when trying to login user")
	}
	if response.StatusCode > 299 {
		var restErr errors.RestError
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("Invalid error interface when trying to login user")
		}
		return nil, &restErr
	}
	var user users.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("Error when trying to unmarshal users login response")
	}
	return &user, nil
}
