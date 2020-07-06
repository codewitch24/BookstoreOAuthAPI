package app

import (
	"github.com/codewitch24/BookstoreOAuthAPI/src/domain/access_token"
	"github.com/codewitch24/BookstoreOAuthAPI/src/http"
	"github.com/codewitch24/BookstoreOAuthAPI/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	handler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access-token/:access_token_id", handler.GetById)
	router.POST("/oauth/access-token", handler.Create)
	_ = router.Run(":8080")
}
