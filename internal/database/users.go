package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func (m *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO users (email, name, password) VALUES ($1, $2, $3) RETURNING id"

	err := m.DB.QueryRowContext(ctx, query, user.Email, user.Name, user.Password).Scan(&user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Get(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	
	query := "SELECT id, email, name, password FROM users WHERE id = $1"

	var user User
	err := m.DB.QueryRowContext(ctx, query, id).Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}
	return &user, nil
}
