package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func Create(data string) (string, error) {
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

	if len(Key) != 32 {
		return "", fmt.Errorf("key value must be 32 bytes. len: %v", len(Key))
	}

	return paseto.NewV2().Encrypt(Key, jsonToken, Footer)
}
