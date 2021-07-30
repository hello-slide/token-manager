package token

import (
	"fmt"
	"time"

	"github.com/o1egl/paseto"
)

func Verify(token string, key []byte) (string, error) {
	var newJsonToken paseto.JSONToken
	var newFooter string
	now := time.Now()

	err := paseto.NewV2().Decrypt(token, key, &newJsonToken, &newFooter)
	if err != nil {
		return "", err
	}

	if newJsonToken.Expiration.Sub(now) <= 0 {
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
