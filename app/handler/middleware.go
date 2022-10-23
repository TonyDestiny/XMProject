package handler

import (
	"net/http"

	"github.com/go-chi/jwtauth"
)

func userIdentity(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := jwtauth.TokenFromHeader(r)
		_, err := jwtauth.VerifyToken(tokenAuth, token)

		if token == "" || err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
