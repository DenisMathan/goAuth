package bakery

import (
	"log"
	"net/http"
	"time"
)

func CreateCookie(w http.ResponseWriter, name string, val string, expiry time.Time) {
	log.Println(val)
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    val,
		Expires:  expiry,
		HttpOnly: true,
		Secure:   true,
		Domain:   "localhost",
		Path:     "/",
	})
}
