package httpHandler

import (
	"log"
	"net/http"

	"github.com/denismathan/goAuth/src/configurations"
	"github.com/denismathan/goAuth/src/database"
	"github.com/denismathan/goAuth/src/router"
	"github.com/denismathan/goAuth/src/services/authentication"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type httpServer struct {
	app *http.Server
	db  database.SqlHandler
	cfg configurations.Config
}

func NewHttpServer() httpServer {
	cfg := configurations.GetConfig()
	sqlHandler := database.NewSqlHandler(cfg.Database)
	oauthConfig := &oauth2.Config{
		ClientID:     cfg.Authorization.ClientID,
		ClientSecret: cfg.Authorization.ClientSecret,
		RedirectURL:  cfg.Authorization.RedirectURL,
		Scopes:       cfg.Authorization.Scopes,
		Endpoint:     google.Endpoint,
	}
	authService := authentication.NewAuthService(&sqlHandler, oauthConfig)
	server := httpServer{
		app: &http.Server{Addr: ":" + cfg.ServerPort, Handler: router.NewRouter(&sqlHandler, &authService)},
		db:  sqlHandler,
		cfg: cfg,
	}
	return server
}

func (server *httpServer) Start() {
	err := server.app.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("Http Server stopped unexpected")
		// s.Shutdown()
	} else {
		log.Println("Http Server stopped")
	}
}
