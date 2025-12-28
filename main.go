package main

import (
	"github/rohitth10e/db"
	user_routes "github/rohitth10e/routes/Users"
	product_routes "github/rohitth10e/routes/products"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.ConnectDB()

	user_routes.Users(server)
	product_routes.Products(server)

	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}
