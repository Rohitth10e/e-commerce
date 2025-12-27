package main

import (
	"github/rohitth10e/db"
	routes "github/rohitth10e/routes/users"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.ConnectDB()

	routes.Users(server)

	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}
