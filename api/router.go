package api

import (
	"database/sql"
	"restaurant/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "restaurant/api/docs"
)

// @title Auth Service
// @version 1.0
// @description This is the Auth service of SitandEat project

// @contact.name Saidakbar
// @contact.url http://www.support_me_with_smile
// @contact.email pardaboyevsaidakbar103@gmail.com

// @host localhost:7777
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
