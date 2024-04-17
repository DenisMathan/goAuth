package middleware

import (
	"context"
	"net/http"
	"strconv"

	"github.com/denismathan/goAuth/src/bakery"
	"github.com/denismathan/goAuth/src/entities"
	"github.com/denismathan/goAuth/src/router/interfaces"
)

type SqlHandler interface {
	FindUser(obj interface{}, email string)
	FindTokenuser(obj interface{}, token string)
	Create(obj interface{})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" || r.URL.Path == "/auth/google/callback" || r.URL.Path == "/auth/google/login" {
			next.ServeHTTP(w, r)
			return
		}
		//get context and cookies
		values := r.Context().Value("values").(map[string]interface{})
		authService := values["authService"].(interfaces.AuthService)
		sqlHandler := values["sqlHandler"].(interfaces.SqlHandler)
		accessTokenCookie, err := r.Cookie("accessToken")
		if err != nil {
			http.Error(w, "no accesstoken", http.StatusUnauthorized)
			return
		}
		accessToken := accessTokenCookie.Value
		if accessToken == "" {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		userCookie, err := r.Cookie("user")
		if err != nil {
			http.Error(w, "no user", http.StatusUnauthorized)
			return
		}

		//check if accessToken is valid
		userID, newAccessToken, err := authService.Validation(accessToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		//check accesstoken-usercombination
		if userCookie.Value != strconv.Itoa(int(userID)) {
			http.Error(w, "invalid accesstoken", http.StatusUnauthorized)
			return
		}

		if newAccessToken != nil {
			sqlHandler.Create(entities.Token{
				Token:          newAccessToken.AccessToken,
				UserID:         userID,
				ExpirationDate: newAccessToken.Expiry,
			})
			bakery.CreateCookie(w, "accessToken", newAccessToken.AccessToken, newAccessToken.Expiry)
			bakery.CreateCookie(w, "user", strconv.Itoa(int(userID)), newAccessToken.Expiry)
		}
		r = r.WithContext(context.WithValue(r.Context(), "userID", userID))
		r = r.WithContext(context.WithValue(r.Context(), "accessToken", accessToken))
		next.ServeHTTP(w, r)
	})
}
