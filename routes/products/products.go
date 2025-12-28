package routes

import (
	handlers "github/rohitth10e/handlers/products"
	"github/rohitth10e/middleware"

	"github.com/gin-gonic/gin"
)

func Products(server *gin.Engine) {
	server.POST("/products/register", middleware.AuthMiddleware(), handlers.RegisterProduct)
	// server.POST("/login", handlers.LoginUser)
}
