package interfaces

import (
	"net/http"

	"golang.org/x/oauth2"
)

type AuthService interface {
	OauthCallback(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Validation(token string) (uint, *oauth2.Token, error)
	// RefreshAccessToken(accessToken string) (*entities.Token, error)
}
