package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int       `json:"id"`
	ProductName string    `json:"product_name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Quantity    int       `json:"quantity"`
	InStock     bool      `json:"in_stock"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
}
