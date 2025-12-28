package handlers

import (
	"fmt"
	"github/rohitth10e/models"
	"github/rohitth10e/repository"
	"github/rohitth10e/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing data",
			"error":   err.Error(),
		})
		return
	}

	if user.Email == "" || user.Password == "" || user.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing required fields",
		})
		return
	}

	hashedPass, err := utils.HashPassword(user.Password)

	if err != nil {
		fmt.Print("[Bcrypt] Something went wrong")
	}

	id, err := repository.InsertUser(user.Name, user.Email, hashedPass, user.Role)
	if err != nil {
		fmt.Println("[DB] Insert error:", err)
		ctx.JSON(http.StatusConflict, gin.H{"success": false, "error": "email already exists"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    gin.H{"id": id, "email": user.Email},
	})
	fmt.Println("[AUTH] User registered with id:", id)
}

func LoginUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "error parsing data",
			"error":   err,
		})
		return
	}

	if user.Email == "" || user.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing required fields",
		})
		return
	}

	dbUser, err := repository.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
		return
	}

	if !utils.VerifyPassword(user.Password, dbUser.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "invalid password"})
		return
	}

	token, err := utils.Sign(user.Email, user.Role)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in", "data": gin.H{
		"id":    dbUser.ID,
		"name":  dbUser.Name,
		"email": dbUser.Email,
		"role":  dbUser.Role,
	}, "token": token})
}
