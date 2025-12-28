package handlers

import (
	"github/rohitth10e/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type req struct {
// 	ID          int       `json:"id"`
// 	reqName string    `json:"req_name"`
// 	Description string    `json:"description"`
// 	Price       float32   `json:"price"`
// 	Quantity    int       `json:"quantity"`
// 	InStock     bool      `json:"in_stock"`
// 	CategoryID  int       `json:"category_id"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

type ProductRequest struct {
	ProductName  string  `json:"product_name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	Quantity     int     `json:"quantity"`
	InStock      bool    `json:"in_stock"`
	CategoryName string  `json:"category_name"`
}

func RegisterProduct(ctx *gin.Context) {
	var req ProductRequest
	err := ctx.ShouldBindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "error parsing data",
			"error":   err.Error(),
		})
		return
	}

	if req.ProductName == "" || req.Description == "" || req.CategoryName == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "missing required fields",
		})
		return
	}

	data, err := repository.InsertProduct(req.ProductName, req.Description, req.Price, req.Quantity, req.InStock, req.CategoryName)
	// req_name string, description string, price float32, quantity int, in_Stock bool, category_name string
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data": gin.H{
			"category": data.Category_name,
			"req":      data.ProductName,
		},
	})
}
