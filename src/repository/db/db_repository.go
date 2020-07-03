package db

import (
	"github.com/codewitch24/BookstoreOAuthAPI/src/clients/cassandra"
	"github.com/codewitch24/BookstoreOAuthAPI/src/domain/access_token"
	"github.com/codewitch24/BookstoreOAuthAPI/src/utils/errors"
)

func NewRepository() DatabaseRepository {
	return &databaseRepository{}
}

type DatabaseRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
}

type databaseRepository struct {
}

func (r *databaseRepository) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()
	return nil, errors.NewInternalServerError("Database connection not implement yet")
}
