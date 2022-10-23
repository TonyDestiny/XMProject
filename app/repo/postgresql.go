package repo

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	HOST = "database"
	PORT = 5432
)

func NewPostgresDB(dbUser, dbPassword, dbName string) (*sqlx.DB, error) {

	dataSource := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		HOST, PORT, dbUser, dbName, dbPassword)
	db, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, err
}
