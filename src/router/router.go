package router

import (
	"context"
	"io"
	"net/http"

	"github.com/denismathan/goAuth/src/router/interfaces"
	"github.com/denismathan/goAuth/src/router/middleware"
	"github.com/gorilla/mux"
)

func NewRouter(sqlH interfaces.SqlHandler, authService interfaces.AuthService) *mux.Router {
	/* Set the router */
	router := mux.NewRouter().StrictSlash(false)
	router.Use(cors)
	router.Use(func(next http.Handler) http.Handler { return fillContext(next, sqlH, authService) })
	//Authentication Validation
	router.Use(middleware.AuthMiddleware)
	requests(router)

	return router
}

func fillContext(next http.Handler, sqlHandler interfaces.SqlHandler, authService interfaces.AuthService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		values := map[string]interface{}{
			"sqlHandler":  sqlHandler,
			"authService": authService,
		}
		newCtx := context.WithValue(r.Context(), "values", values)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerControl := "Authorization, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Methods, Access-Control-Request-Headers"
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Headers", headerControl)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func requests(router *mux.Router) {
	router.HandleFunc("/", routeBase)
	router.HandleFunc("/auth/google/login", login)
	router.HandleFunc("/auth/google/callback", oauthGoogleCallback)
	todoRequests(router)
	// tokenRequests(router)
	// userRequests(router)
}

func routeBase(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	io.WriteString(writer, "hihihi")
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Methods, Access-Control-Request-Headers")
}
