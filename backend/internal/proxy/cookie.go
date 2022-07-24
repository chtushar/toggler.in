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
	R 		*http.Request
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

	cookie := &http.Cookie{
			Name:     "auth",
			Value:    cookieCoded,
			Path:     "/",
			HttpOnly: true,
	}

	cfg.R.AddCookie(cookie)

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

func ValidateAuthCookie(w *http.ResponseWriter, r *http.Request, sc *securecookie.SecureCookie, jwt *helpers.JWT) error {

	cookie, err := r.Cookie("auth");


	if err != nil {
		return err
	}

	var decodedCookieValue string;
	err = sc.Decode("auth", cookie.Value, &decodedCookieValue)

	if err != nil {
		return err
	}


	_, err = jwt.ReadTokenAndValidate(decodedCookieValue)

	if err != nil  {
		return err
	}

	return nil
}