package proxy

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"toggler.in/internal/helpers"
)

type AuthCookieConfig struct {
	JWT 	*helpers.JWT
	SC    *securecookie.SecureCookie
	W 		*http.ResponseWriter
	User 	map[string]interface{}
}


func NewSecureCookie(hashKey []byte, blockKey []byte) *securecookie.SecureCookie {
	return securecookie.New(hashKey, blockKey)
}

func SetAuthCookie(cfg *AuthCookieConfig) error {
	token, err := cfg.JWT.NewToken(helpers.AuthSecret, cfg.User)

	if err != nil {
		return err
	}

	cookieCoded, err := cfg.SC.Encode("auth", token)

	if err != nil {
		return err
	}

	http.SetCookie(*cfg.W, &http.Cookie{
			Name:     "auth",
			Value:    cookieCoded,
			Path:     "/",
			HttpOnly: true,
		})

	return nil
}

func ClearAuthCookie(w *http.ResponseWriter) error {
	c := &http.Cookie{
    Name:     "auth",
    Value:    "",
    Path:     "/",
    Expires: time.Unix(0, 0),

    HttpOnly: true,
	}

	http.SetCookie(*w, c)

	return nil
}