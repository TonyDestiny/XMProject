package user

import (
	"crypto/sha1"
	"fmt"
)

const (
	salt = "djnzlxnlsd"
)

type UsersStore interface {
	CreateUser(user User) (int, error)
	GetUser(username, password string) (User, error)
}

type UsersService struct {
	ustore UsersStore
}

func NewUsersService(ustore UsersStore) *UsersService {
	return &UsersService{ustore: ustore}
}

func (u UsersService) CreateUser(user User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return u.ustore.CreateUser(user)
}

func (u UsersService) GetUser(username, password string) (User, error) {
	user, err := u.ustore.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
