package authentication

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/denismathan/goAuth/src/entities"
	"golang.org/x/oauth2"
)

func (authService *AuthService) Validation(accessToken string) (uint, *oauth2.Token, error) {
	tokenUser := authService.getTokenUser(accessToken)
	log.Println(tokenUser)
	var newAccessToken *oauth2.Token
	//checks if token is currently stored in DB
	if authService.CheckExpiracy(tokenUser) {
		expDate := tokenUser.ExpirationDate.Unix()
		now := time.Now().Unix()
		if expDate >= now && expDate-(60*5) < now {
			log.Println("new token")
			newAccessToken, _ = authService.refreshAccessToken(tokenUser.UserID)
		}
		return tokenUser.UserID, newAccessToken, nil
	}
	return 0, newAccessToken, errors.New("token expired")
}

func (authService *AuthService) refreshAccessToken(userID uint) (*oauth2.Token, error) {
	user := entities.User{}
	authService.sqlHandler.FindUserById(&user, userID)
	// Erstelle ein TokenSource mit dem Config und dem Refresh-Token
	tokenSource := authService.oauthConfig.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: user.RefreshToken,
	})

	// Hole den neuen AccessToken
	token, err := tokenSource.Token()
	if err != nil {
		return nil, err
	}

	// Gib den neuen AccessToken zurück
	return token, nil
}

func (authService *AuthService) CheckExpiracy(tokenUser entities.Token) bool {
	log.Println(time.Time(tokenUser.ExpirationDate))
	log.Println(time.Now())
	return tokenUser.ExpirationDate.Unix() >= time.Now().Unix()
}

func (authService *AuthService) getTokenUser(accessToken string) entities.Token {
	tokenUser := entities.Token{}
	authService.sqlHandler.FindTokenuser(&tokenUser, accessToken)
	return tokenUser
}
