package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.0.1:8080/users/login",
		ReqBody:      `{"email":"john@gmail.com","password":"secret"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("john@gmail.com", "secret")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.0.1:8080/users/login",
		ReqBody:      `{"email":"john@gmail.com","password":"secret"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": "404", "error": "NOT_FOUND"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("john@gmail.com", "secret")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.0.1:8080/users/login",
		ReqBody:      `{"email":"john@gmail.com","password":"secret"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Invalid login credentials", "status": 404, "error": "NOT_FOUND"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("john@gmail.com", "secret")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "Invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.0.1:8080/users/login",
		ReqBody:      `{"email":"john@gmail.com","password":"secret"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1", "first_name": "Jimmy", "last_name": "Ford" , "email": "jimmy@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("john@gmail.com", "secret")
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Error when trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://127.0.0.1:8080/users/login",
		ReqBody:      `{"email":"john@gmail.com","password":"secret"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1, "first_name": "Jimmy", "last_name": "Ford" , "email": "jimmy@gmail.com"}`,
	})
	repository := usersRepository{}
	user, err := repository.LoginUser("john@gmail.com", "secret")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "Jimmy", user.FirstName)
	assert.EqualValues(t, "Ford", user.LastName)
	assert.EqualValues(t, "jimmy@gmail.com", user.Email)
}
