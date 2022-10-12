package infrastructure

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/4strodev/todo_app/pkg/todo/user/domain"
)

type SqliteUserRepository struct {
	pool *sql.DB
}

func (self *SqliteUserRepository) SetPool(pool *sql.DB) {
	self.pool = pool
}

// Get all users
func (self *SqliteUserRepository) GetUsers() ([]domain.User, error) {
	users := []domain.User{}

	rows, err := self.pool.Query("SELECT id, username, password FROM users")
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return users, errors.New("Error getting users")
	}

	var user domain.User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}

	return users, nil
}

// Get user based on id
func (self *SqliteUserRepository) GetUser(id string) (domain.User, error) {
	user := domain.User{}

	rows, err := self.pool.Query("SELECT id, username, password FROM users WHERE id = ?", id)

	defer rows.Close()
	if err != nil {
		log.Println(err)
		return user, fmt.Errorf("Error getting user identified with %s", id)
	}

	for rows.Next() {
		rows.Scan(&user.Id, &user.Username, &user.Password)
	}

	return user, nil
}

// Save user to sqlite
func (self *SqliteUserRepository) SaveUser(user domain.User) error {
	insertQuery := "INSERT INTO users (id, username, password) VALUES (?, ?, ?)"

	rows, err := self.pool.Query(insertQuery, user.Id, user.Username, user.Password)
	defer rows.Close()
	if err != nil {
		log.Println(err)
		return errors.New("Error saving new user")
	}

	return nil
}

// Update user in sqlite
func (self *SqliteUserRepository) UpdateUser(user domain.User) (domain.User, error) {
	updateQuery := "UPDATE users SET username = ?, password = ? WHERE id = ?"

	rows, err := self.pool.Query(updateQuery, user.Username, user.Password, user.Id)
	rows.Close()
	if err != nil {
		log.Println(err)
		return user, fmt.Errorf("Error updating user identified with %s", user.Id)
	}

	return user, nil
}

// Delete user from sqlite
func (self *SqliteUserRepository) DeleteUser(id string) error {
	deleteQuery := "DELETE FROM users WHERE id = ?"

	rows, err := self.pool.Query(deleteQuery, id)
	rows.Close()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Error updating user identified with %s", id)
	}

	return nil
}
