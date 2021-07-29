package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func Verify(token string) (string, error) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	now := time.Now()

	err := paseto.NewV2().Decrypt(token, Key, &newJsonToken, &newFooter)
	if err != nil {
		return "", err
	}

	if now.Sub(newJsonToken.Expiration) <= 0 {
		return "", fmt.Errorf("The token has expired.")
	}

	if newJsonToken.Audience != Audience ||
		newJsonToken.Issuer != Issuer ||
		newJsonToken.Jti != Jti ||
		newFooter != Footer {
		return "", fmt.Errorf("The token information is incorrect.")
	}

	return newJsonToken.Get("data"), nil
}
