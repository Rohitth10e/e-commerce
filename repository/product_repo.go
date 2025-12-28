package repository

import (
	"context"
	"github/rohitth10e/db"
	"time"
)

type Data struct {
	ID            int       `json:"id"`
	ProductName   string    `json:"product_name"`
	Description   string    `json:"description"`
	Price         float32   `json:"price"`
	Quantity      int       `json:"quantity"`
	InStock       bool      `json:"in_stock"`
	CategoryID    int       `json:"category_id"`
	Category_name string    `json:"category_name"`
	CreatedAt     time.Time `json:"created_at"`
}

func InsertProduct(product_name string, description string, price float32, quantity int, in_Stock bool, category_name string) (Data, error) {
	// query := `INSERT INTO PRODUCT (product_name,description,price, quantity, in_Stock) VALUES ($1, $2, $3, $4, $5)`
	var CategoryID int
	err := db.Db.QueryRowContext(context.Background(),
		`INSERT INTO category (category_name)
         VALUES ($1)
         ON CONFLICT (category_name) DO UPDATE SET category_name = EXCLUDED.category_name
         RETURNING id`,
		category_name).Scan(&CategoryID)

	if err != nil {
		return Data{}, err
	}

	var data Data
	err = db.Db.QueryRowContext(context.Background(),
		`INSERT INTO product (product_name, description, price, quantity, in_stock, category_id, created_at)
         VALUES ($1, $2, $3, $4, $5, $6, NOW())
         RETURNING id, product_name, description, price, quantity, in_stock, category_id, created_at`,
		product_name, description, price, quantity, in_Stock, CategoryID).
		Scan(&data.ID, &data.ProductName, &data.Description, &data.Price,
			&data.Quantity, &data.InStock, &data.CategoryID, &data.CreatedAt)

	if err != nil {
		return Data{}, err
	}

	data.Category_name = category_name

	return data, nil
}
