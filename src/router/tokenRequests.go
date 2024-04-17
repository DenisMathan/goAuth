package router

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/denismathan/goAuth/src/router/interfaces"
)

func login(w http.ResponseWriter, r *http.Request) {
	authService := r.Context().Value("values").(map[string]interface{})["authService"].(interfaces.AuthService)
	authService.Login(w, r)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	return state
}

func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	authService := r.Context().Value("values").(map[string]interface{})["authService"].(interfaces.AuthService)
	authService.OauthCallback(w, r)
}

// func tokenRequests(router *mux.Router) {
// 	// router.HandleFunc("/api/createTodo", createTodo)
// 	router.HandleFunc("/refreshToken", refreshAccessToken).Methods("GET", "OPTIONS")
// }

// func refreshAccessToken(w http.ResponseWriter, r *http.Request) {
// 	authService := r.Context().Value("values").(map[string]interface{})["authService"].(interfaces.AuthService)
// 	accessToken := r.Context().Value("accessToken").(string)
// 	authService.RefreshAccessToken(accessToken)
// }

// func readTodos(w http.ResponseWriter, request *http.Request) {
// 	sqlH := request.Context().Value("values").(map[string]interface{})["sqlHandler"].(interfaces.SqlHandler)
// 	userID := request.Context().Value("userId").(uint)
// 	var todos []entities.Todo
// 	sqlH.GetTodos(&todos, userID)
// 	json.NewEncoder(w).Encode(todos)
// }
