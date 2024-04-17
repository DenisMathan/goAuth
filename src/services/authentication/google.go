package authentication

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/denismathan/goAuth/src/entities"
)

type googleResp struct {
	Email          string `json:"email"`
	ExpirationDate string `json:"exp"`
	EmailVerified  string `json:"email_verified"`
}

type googleUser struct {
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	UserName   string `json:"name"`
	Name       string `json:"given_name"`
	Surname    string `json:"family_name"`
	ProfilePic string `json:"picture"`
	Locale     string `json:"locale"`
}

func (googleAuthService *AuthService) googleRespToTokenUser(resp googleResp) entities.Token {
	unixTime, _ := strconv.ParseInt(resp.ExpirationDate, 10, 64)
	return entities.Token{
		ExpirationDate: time.Unix(unixTime, 0),
		UserEmail:      resp.Email,
	}
}

func (googleAuthService *AuthService) googleRespToUser(resp []byte) entities.User {
	respStruct := googleUser{}
	err := json.Unmarshal(resp, &respStruct)
	if err != nil {
		log.Println(err)
	}

	return entities.User{
		Email:    respStruct.Email,
		UserName: respStruct.UserName,
		Name:     respStruct.Name,
		Surame:   respStruct.Surname,
		Locale:   respStruct.Locale,
		Verified: respStruct.Verified,
	}
}
