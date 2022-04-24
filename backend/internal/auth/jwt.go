package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"toggler.in/internal/helpers"
)

type JWT struct {
	secret string
}

type tokenType struct {
	TypeName          string
	ExpirationMinutes int
}

var AuthSecret = &tokenType{
	TypeName:          "auth",
	ExpirationMinutes: 365 * 24 * 60,
}

const (
	KeyIssuer = "key"
	KeyIssuedAt = "iat"
	KeyExpiry = "exp"
	KeyTokenType = "type"
	KeyUserId = "id"
	KeyUserEmail = "email"
	KeyUserName = "name"
)

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) NewToken(tokenType *tokenType, m map[string]interface{}) (string, error) {
	claim := buildClaim(tokenType, m)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(j.secret))
}

func defaultClaim() jwt.MapClaims {
	return jwt.MapClaims{
		KeyIssuer: "toggler.in",
		KeyIssuedAt:  time.Now().Unix(),
	}
}

func buildClaim(tokenType *tokenType, m map[string]interface{}) jwt.MapClaims {
	var mapClaim = defaultClaim()

	for key, value := range m {
		mapClaim[key] = value
	}

	mapClaim[KeyTokenType] = tokenType.TypeName
	mapClaim[KeyExpiry] = helpers.UNIXTimestampFromNow(tokenType.ExpirationMinutes)

	return mapClaim
}
