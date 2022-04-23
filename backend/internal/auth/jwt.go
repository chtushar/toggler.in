package auth

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}