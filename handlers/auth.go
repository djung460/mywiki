package handlers

import (
	"net/http"

	"gopkg.in/authboss.v1"
)

// HandleAuthboss handles all things authboss
func HandleAuthboss(ab *authboss.Authboss) http.Handler {
	return ab.NewRouter()
}
