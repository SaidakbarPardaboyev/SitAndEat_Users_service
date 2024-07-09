package api

import (
	"database/sql"
	"restaurant/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	auth := r.Group("/users")

	h := handler.NewHandler(db)
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	return r
}
