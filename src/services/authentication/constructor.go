package authentication

import (
	"golang.org/x/oauth2"
)

type SqlHandler interface {
	FindUser(obj interface{}, email string)
	FindUserById(obj interface{}, id uint)
	FindTokenuser(obj interface{}, token string)
	Create(obj interface{})
	DeleteExpiredTokens(userID uint)
}

type AuthService struct {
	oauthConfig *oauth2.Config
	sqlHandler  SqlHandler
}

type LoginUrl struct {
	URL string
}

func NewAuthService(sqlHandler SqlHandler, config *oauth2.Config) AuthService {
	authService := new(AuthService)
	authService.oauthConfig = config
	authService.sqlHandler = sqlHandler
	return *authService
}
