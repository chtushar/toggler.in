package helpers

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
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

var (
	ErrorSigningKey = errors.New("unexpected signing method")
	ErrorTypeCast   = errors.New("error typecasting token claim")
)


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

func (j *JWT) jwtKey(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorSigningKey
		}
		return []byte(j.secret), nil
	}

// // ReadToken reads and validates a signed token
func (j *JWT)ReadToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, j.jwtKey)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (j *JWT)ReadTokenAndValidate(tokenStr string) (*jwt.Token, error) {
	token, err := j.ReadToken(tokenStr)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
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
	mapClaim[KeyExpiry] = UNIXTimestampFromNow(tokenType.ExpirationMinutes)

	return mapClaim
}
