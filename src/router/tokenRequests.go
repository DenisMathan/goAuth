package router

import (
	"net/http"

	"github.com/denismathan/goAuth/src/router/interfaces"
)

func login(w http.ResponseWriter, r *http.Request) {
	authService := r.Context().Value("values").(map[string]interface{})["authService"].(interfaces.AuthService)
	authService.Login(w, r)
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	authService := r.Context().Value("values").(map[string]interface{})["authService"].(interfaces.AuthService)
	authService.OauthCallback(w, r)
}
