package token

import (
	"time"

	"github.com/o1egl/paseto"
)

func Create(data string, key []byte) (string, error) {
	now := time.Now()
	exp := now.Add(24 * time.Hour)
	nbt := now

	jsonToken := paseto.JSONToken{
		Audience:   "token-manager-hello-slide",
		Issuer:     "token-manager",
		Jti:        "glpcsdd",
		Subject:    "settion-token",
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  nbt,
	}
	jsonToken.Set("data", data)

	return paseto.NewV2().Encrypt(key, jsonToken, Footer)
}
