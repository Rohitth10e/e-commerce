package main

import (
	"github/rohitth10e/db"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	db.ConnectDB()

	if err := server.Run(":3000"); err != nil {
		panic(err)
	}
}
