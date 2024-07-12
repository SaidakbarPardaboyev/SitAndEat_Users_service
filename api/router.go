package api

import (
	"database/sql"
	"restaurant/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "restaurant/api/docs"
)

// @title           Foydalanuvchilarni boshqarish APIsi
// @version         1.0
// @description     Bu foydalanuvchi boshqaruvi uchun namuna serveri.

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0

// @host      localhost:7777
// @schemes http
func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/users")
 
	h := handler.NewHandler(db)
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)

	return r
}
