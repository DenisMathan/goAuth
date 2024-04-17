package authentication

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/denismathan/goAuth/src/bakery"
	"golang.org/x/oauth2"
)

func (authService *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	oauthState := generateStateOauthCookie()
	log.Println(oauthState)
	u := authService.oauthConfig.AuthCodeURL(oauthState, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
	login := LoginUrl{
		URL: u,
	}
	json.NewEncoder(w).Encode(login)
}

func generateStateOauthCookie() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

func (authService *AuthService) OauthCallback(w http.ResponseWriter, r *http.Request) {
	userData, token, err := authService.getUserData(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "http://localhost:4200/login", http.StatusTemporaryRedirect)
		return
	}

	userID, err := authService.userLogin(token, userData)

	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "http://localhost:4200/login", http.StatusPermanentRedirect)
		return
	}
	log.Println(token.AccessToken)
	bakery.CreateCookie(w, "accessToken", token.AccessToken, token.Expiry)
	bakery.CreateCookie(w, "user", strconv.Itoa(int(userID)), token.Expiry)

	http.Redirect(w, r, "http://localhost:4200/home", http.StatusPermanentRedirect)
}

func (authService *AuthService) getUserData(code string) ([]byte, *oauth2.Token, error) {
	// Use code to get token and get user info from Google.
	// tok := oauth2.Token{}
	token, err := authService.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	// authService.Validation(token.AccessToken)
	return contents, token, nil
}
