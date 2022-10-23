package handler

import (
	"XMProject/app/service/company"
	"XMProject/app/service/user"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	serviceCompany company.CompaniesStore
	serviceUser    user.UsersStore
}

func NewHandler(serviceCompany company.CompaniesStore, serviceUser user.UsersStore) *Handler {
	return &Handler{
		serviceCompany: serviceCompany,
		serviceUser:    serviceUser,
	}
}

func (h *Handler) InitRoutes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Post("/sing-up", h.signUp)
	})

	r.Route("/company", func(r chi.Router) {

		r.With(userIdentity).Post("/add", h.createCompany)

		r.Get("/get", h.getCompany)

		r.With(userIdentity).Delete("/del", h.delCompany)

		r.With(userIdentity).Patch("/patch", h.updateCompany)
	})

	return r
}
