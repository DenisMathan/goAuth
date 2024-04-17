package configurations

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database      Database
	ServerPort    string `envconfig:"SERVER_PORT" default:"80"`
	Authorization Authorization
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" required:"true"`
	Port     int    `envconfig:"DATABASE_PORT" required:"true"`
	User     string `envconfig:"DATABASE_USER" required:"true"`
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

type Authorization struct {
	ClientID     string   `envconfig:"GOOGLE_CLIENT_ID" required:"true"`
	ClientSecret string   `envconfig:"GOOGLE_CLIENT_SECRET" required:"true"`
	RedirectURL  string   `envconfig:"GOOGLE_REDIRECT_URL" required:"true"`
	Scopes       []string `envconfig:"GOOGLE_SCOPES" required:"true"`
}

func NewParsedConfig() Config {
	_ = godotenv.Load(".env")
	cnf := Config{}
	envconfig.Process("", &cnf)
	return cnf
}

func GetConfig() Config {
	_ = godotenv.Load(".env")
	cnf := Config{}
	envconfig.Process("", &cnf)

	cnf.Authorization.Scopes = strings.Split(cnf.Authorization.Scopes[0], " ")
	return cnf
}
