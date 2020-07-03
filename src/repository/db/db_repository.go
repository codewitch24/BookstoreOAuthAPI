package db

import (
	"github.com/codewitch24/BookstoreOAuthAPI/src/clients/cassandra"
	"github.com/codewitch24/BookstoreOAuthAPI/src/domain/access_token"
	"github.com/codewitch24/BookstoreOAuthAPI/src/utils/errors"
	"github.com/gocql/gocql"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires = ? WHERE access_token = ?;"
)

func NewRepository() DatabaseRepository {
	return &databaseRepository{}
}

type DatabaseRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type databaseRepository struct {
}

func (r *databaseRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestError) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenId).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token found with given id")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}
	return &result, nil
}

func (r *databaseRepository) Create(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		at.AccessToken,
		at.UserId,
		at.ClientId,
		at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *databaseRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		at.Expires,
		at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
