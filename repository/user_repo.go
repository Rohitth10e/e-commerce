package repository

import (
	"context"
	"fmt"
	"github/rohitth10e/db"
)

func InsertUser(name, email, password, role string) (int, error) {
	query := `INSERT INTO users (name, email, password, role, created_at)
	          VALUES ($1, $2, $3, $4, NOW()) RETURNING id`

	var id int
	err := db.Db.QueryRowContext(context.Background(), query, name, email, password, role).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insert failed: %w", err)
	}
	return id, nil
}
