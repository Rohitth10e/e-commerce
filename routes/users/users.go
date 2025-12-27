package routes

import (
	handlers "github/rohitth10e/handlers/users"

	"github.com/gin-gonic/gin"
)

func Users(server *gin.Engine) {
	server.POST("/register", handlers.RegisterUser)
}
