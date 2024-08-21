package api

import (
	"database/sql"
	"restaurant/api/handler"

	_ "restaurant/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth Service
// @version 1.0
// @description This is the Auth service of SitandEat project


// @BasePath /users

func NewRouter(db *sql.DB) *gin.Engine {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/users")

	h := handler.NewHandler(db)
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	return r
}
