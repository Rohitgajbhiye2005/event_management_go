package database

import (
	"context"
	"database/sql"
	"fmt"
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

// func (m *UserModel) Insert(user *User) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	query := "INSERT INTO users (email, password, name) VALUES(?, ?, ?)"

// 	// Use ExecContext instead of QueryRowContext for SQLite
// 	result, err := m.DB.ExecContext(ctx, query, user.Email, user.Password, user.Name)
// 	if err != nil {
// 		return err
// 	}

// 	// Get the last insert ID separately
// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return err
// 	}

// 	user.Id = int(id)
// 	return nil
// }

func (m *UserModel) Insert(user *User) error {
	fmt.Printf("DEBUG: Inserting user - Email: %s, Name: %s\n", user.Email, user.Name) // ADD THIS
	
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "INSERT INTO users (email, password, name) VALUES(?, ?, ?)"
	fmt.Printf("DEBUG: Query: %s\n", query) // ADD THIS

	// Use ExecContext instead of QueryRowContext for SQLite
	result, err := m.DB.ExecContext(ctx, query, user.Email, user.Password, user.Name)
	if err != nil {
		fmt.Printf("DEBUG: ExecContext ERROR: %v\n", err) // ADD THIS
		return err
	}

	// Get the last insert ID separately
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("DEBUG: LastInsertId ERROR: %v\n", err) // ADD THIS
		return err
	}

	user.Id = int(id)
	fmt.Printf("DEBUG: Success! User ID: %d\n", user.Id) // ADD THIS
	return nil
}