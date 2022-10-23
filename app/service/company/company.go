package company

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type CompaniesStore interface {
	Create(c *Company) (uuid.UUID, error)
	Get(nameCompany string) (*Company, error)
	Patch(c *Company) error
	Delete(nameCompany string) error
}

type CompaniesService struct {
	cstore CompaniesStore
}

func NewCompaniesService(cstore CompaniesStore) *CompaniesService {
	return &CompaniesService{
		cstore: cstore,
	}
}

func (cs *CompaniesService) Create(c *Company) (uuid.UUID, error) {
	c.ID = uuid.New()

	if valErr := validateCompanyName(c.Name); valErr != nil {
		return uuid.Nil, valErr
	}

	u, err := cs.cstore.Create(c)
	if err != nil {
		return uuid.Nil, err
	}

	return u, nil
}

func (cs *CompaniesService) Get(nameCompany string) (*Company, error) {
	return cs.cstore.Get(nameCompany)
}

func (cs *CompaniesService) Patch(c *Company) error {
	return cs.cstore.Patch(c)
}

func (cs *CompaniesService) Delete(nameCompany string) error {
	return cs.cstore.Delete(nameCompany)
}

func validateCompanyName(companyName string) error {
	r := []rune(companyName)
	if len(r) > 15 {
		return errors.New("company name must be no more than 15 characters")
	}

	if splitBySpace := strings.Split(companyName, " "); len(splitBySpace) > 1 {
		return fmt.Errorf("company name has whitespace. Please replace the spaces with the symbol \"_\". Example: %s",
			strings.Join(splitBySpace, "_"))
	}

	return nil
}
