package proxy

import "github.com/gorilla/securecookie"


func NewSecureCookie(hashKey []byte, blockKey []byte) *securecookie.SecureCookie {
	return securecookie.New(hashKey, blockKey)
}
