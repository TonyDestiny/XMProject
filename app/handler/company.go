package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"XMProject/app/service/company"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createCompany(_ http.ResponseWriter, r *http.Request) {
	var c company.Company

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error decode post body: %v", err))
	}

	id, errUser := h.serviceCompany.Create(&c)
	if errUser != nil {
		newErrorResponse(fmt.Sprintf("error create company: %v", errUser))
	} else {
		logrus.Infof("Add company %s to DB with ID=%s", c.Name, id.String())
	}
}

func (h *Handler) getCompany(_ http.ResponseWriter, r *http.Request) {
	companyName := r.URL.Query().Get("company")

	res, err := h.serviceCompany.Get(companyName)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error get company by name %s: %v", companyName, err))
	} else {
		logrus.Infof("Get company %s: %v", companyName, res)
	}
}

func (h *Handler) delCompany(_ http.ResponseWriter, r *http.Request) {
	companyName := r.URL.Query().Get("company")

	err := h.serviceCompany.Delete(companyName)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error delete company by name %s: %v", companyName, err))
	} else {
		logrus.Infof("Delete company %s", companyName)
	}
}

func (h *Handler) updateCompany(_ http.ResponseWriter, r *http.Request) {
	var c company.Company

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error decode patch body: %v", err))
	}

	if valErr := c.IsValid(); err != nil {
		newErrorResponse(fmt.Sprintf("incorrect patch data: %v", valErr))
	}

	err = h.serviceCompany.Patch(&c)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error patch company %s: %v", c.Name, err))
	} else {
		logrus.Infof("Update company %s: %v", c.Name, c)
	}
}
