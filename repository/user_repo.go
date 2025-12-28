package repository

import (
	"context"
	"fmt"
	"github/rohitth10e/db"
	"github/rohitth10e/models"
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

func GetUserByEmail(email string) (models.User, error) {

	query := `SELECT id, name, email, password, role, created_at FROM USERS WHERE email=($1)`

	var user models.User
	err := db.Db.QueryRowContext(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	fmt.Print(user)

	if err != nil {
		return models.User{}, fmt.Errorf("insert failed: %w", err)
	}

	return user, nil
}
