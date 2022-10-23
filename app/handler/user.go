package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"XMProject/app/service/user"
	"github.com/go-chi/jwtauth"
	"github.com/sirupsen/logrus"
)

const secret = "<jwt-secret>" // Replace <jwt-secret> with your secret key that is private to you.

var tokenAuth = jwtauth.New("HS256", []byte(secret), nil)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var u user.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error decode post body: %v", err))
	}

	if u.Username == "" || u.Password == "" {
		http.Error(w, "missing user or password", http.StatusBadRequest)
		return
	}

	id, err := h.serviceUser.CreateUser(u)
	if err != nil {
		newErrorResponse(fmt.Sprintf("error create user: %v", err))
	} else {
		logrus.Infof("Add user %s to DB with ID=%d", u.Name, id)

		token := makeToken(id)

		http.SetCookie(w, &http.Cookie{
			HttpOnly: true,
			Expires:  time.Now().Add(7 * 24 * time.Hour),
			SameSite: http.SameSiteLaxMode,
			// Uncomment below for HTTPS:
			// Secure: true,
			Name:  "jwt",
			Value: token,
		})

		logrus.Infof("Add user token=%s", token)

	}

}

func makeToken(id int) string {
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"id": id})
	return tokenString
}
