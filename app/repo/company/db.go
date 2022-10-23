package company

import (
	"XMProject/app/service/company"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	tableName = "companies"
)

const (
	queryInsert = `INSERT INTO %s (id, name, description, amount_of_employees, registered, type) values ($1, $2, $3, $4, $5, $6) RETURNING id`
	queryGet    = `SELECT * FROM %s AS c WHERE c.name=$1`
	queryDelete = `DELETE FROM %s WHERE name=$1`
	queryUpdate = `UPDATE %s SET name=$1, description=$2, amount_of_employees=$3, registered=$4, type=$5 WHERE id=$6 RETURNING name`
)

var _ company.CompaniesStore = &RepoCompanies{}

type RepoCompanies struct {
	db *sqlx.DB
}

func NewRepoCompanies(db *sqlx.DB) *RepoCompanies {
	return &RepoCompanies{db: db}
}

func (cs *RepoCompanies) Create(c *company.Company) (uuid.UUID, error) {

	query := fmt.Sprintf(queryInsert, tableName)
	row := cs.db.QueryRow(query, c.ID, c.Name, c.Description, c.AmountOfEmployees, c.Registered, c.Type)

	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (cs *RepoCompanies) Get(nameCompany string) (*company.Company, error) {
	var company company.Company
	query := fmt.Sprintf(queryGet, tableName)

	err := cs.db.Get(&company, query, nameCompany)
	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (cs *RepoCompanies) Patch(c *company.Company) error {
	query := fmt.Sprintf(queryUpdate, tableName)
	row := cs.db.QueryRow(query, c.Name, c.Description, c.AmountOfEmployees, c.Registered, c.Type, c.ID)
	if row.Err() != nil {
		return row.Err()
	}
	return nil
}

func (cs *RepoCompanies) Delete(nameCompany string) error {
	query := fmt.Sprintf(queryDelete, tableName)

	_, err := cs.db.Exec(query, nameCompany)
	if err != nil {
		return err
	}

	return nil
}
