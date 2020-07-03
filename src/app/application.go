package app

import (
	"github.com/codewitch24/BookstoreOAuthAPI/src/clients/cassandra"
	"github.com/codewitch24/BookstoreOAuthAPI/src/domain/access_token"
	"github.com/codewitch24/BookstoreOAuthAPI/src/http"
	"github.com/codewitch24/BookstoreOAuthAPI/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	session.Close()
	repository := db.NewRepository()
	service := access_token.NewService(repository)
	handler := http.NewHandler(service)
	router.GET("/oauth/access-token/:access_token_id", handler.GetById)
	_ = router.Run(":8080")
}
