package company

import (
	"errors"
	"github.com/google/uuid"
)

type TypeCompany int8

const (
	TypeCompany_Invalid TypeCompany = iota
	TypeCompany_NonProfit
	TypeCompany_Cooperative
	TypeCompany_SoleProprietorship
)

var TypeCompany_name = map[int8]string{
	0: "Invalid",
	1: "NonProfit",
	2: "Cooperative",
	4: "SoleProprietorship",
}

var TypeCompany_num = map[string]int8{
	"Invalid":            0,
	"NonProfit":          1,
	"Cooperative":        2,
	"SoleProprietorship": 3,
}

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Company struct {
	ID                uuid.UUID `json:"id" db:"id"`
	Name              string    `json:"name" db:"name"`
	Description       string    `json:"description" db:"description"`
	AmountOfEmployees int32     `json:"amount_of_employees" db:"amount_of_employees"`
	Registered        bool      `json:"registered" db:"registered"`
	Type              string    `json:"type" db:"type"`
}

func (c Company) IsValid() error {
	if !c.validName() {
		return errors.New("incorrect company name")
	}
	if !c.validID() {
		return errors.New("incorrect company ID")
	}
	if !c.validType() {
		return errors.New("incorrect company type")
	}

	return nil
}

func (c Company) validName() bool {
	r := []rune(c.Name)
	return len(r) <= 15
}

func (c Company) validID() bool {
	return c.ID.String() != ""
}

func (c Company) validType() bool {
	if _, ok := TypeCompany_num[c.Type]; ok {
		return true
	}

	return false
}
