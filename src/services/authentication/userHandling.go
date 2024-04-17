package authentication

import (
	"errors"
	"log"

	"github.com/denismathan/goAuth/src/entities"
	"golang.org/x/oauth2"
)

func (service *AuthService) userLogin(token *oauth2.Token, contents []byte) (uint, error) {
	log.Println("userLogin")
	//create UserObject

	//TODO Switch for different authServices
	user := service.googleRespToUser(contents)
	user.RefreshToken = token.RefreshToken
	service.findOrCreateUser(&user)
	if !user.Verified {
		return 0, errors.New("email not verified")
	}
	tokenUser := entities.Token{}

	//find Or Create Token
	service.sqlHandler.FindTokenuser(&tokenUser, token.AccessToken)
	if tokenUser.UserID == 0 {
		service.sqlHandler.DeleteExpiredTokens(user.ID)
		service.sqlHandler.Create(entities.Token{
			Token:          token.AccessToken,
			UserID:         user.ID,
			ExpirationDate: token.Expiry,
		})
		log.Println("token was created")
	}
	return user.ID, nil
}

func (service *AuthService) findOrCreateUser(user *entities.User) error {
	service.sqlHandler.FindUser(&user, user.Email)
	if user.ID == 0 {
		service.sqlHandler.Create(&user)
	}
	return nil
}
