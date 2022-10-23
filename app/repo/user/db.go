package user

import (
	"XMProject/app/service/user"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var _ user.UsersStore = &RepoUsers{}

const (
	usersTable = "users"
)

type RepoUsers struct {
	db *sqlx.DB
}

func NewRepoUsers(db *sqlx.DB) *RepoUsers {
	return &RepoUsers{db: db}
}

func (rs *RepoUsers) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := rs.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (rs *RepoUsers) GetUser(username, password string) (user.User, error) {
	var u user.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := rs.db.Get(&u, query, username, password)

	return u, err
}
