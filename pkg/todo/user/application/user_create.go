package application

import (
	"errors"
	"fmt"
	"log"

	"github.com/4strodev/todo_app/pkg/todo/user/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const USER_SUFIX = "USR"

type UserCreate struct {
	repository domain.UserRepository
}

func (self *UserCreate) Run(username string, password string) error {
	user := domain.User{}

	user.Id = newUserId()
	user.Username = username

	// Encrypting password
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = string(encryptedPassword)

	if err != nil {
		log.Println(err)
		return errors.New("Error encrypting user password")
	}

	return self.repository.SaveUser(user)
}

func newUserId() string {
	return fmt.Sprintf("%s:%s", uuid.NewString(), USER_SUFIX)
}
