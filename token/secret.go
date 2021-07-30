package token

import (
	"context"
	"fmt"

	"github.com/dapr/go-sdk/client"
)

const (
	Audience    = "token-manager-hello-slide"
	Issuer      = "token-manager"
	Jti         = "glpcsdd"
	Subject     = "settion-token"
	Footer      = "This is a Hello slide api session token."
	SecretStore = "google-secret-state"
	KeySecret   = "token-common-key"
)

var Key []byte

func GetKey(client *client.Client, ctx *context.Context) error {
	opt := map[string]string{
		"version": "2",
	}
	secret, err := (*client).GetSecret(*ctx, SecretStore, KeySecret, opt)
	if err != nil {
		return err
	}
	getKey := []byte(secret[KeySecret])
	if len(getKey) != 32 {
		return fmt.Errorf("key value must be 32 bytes.")
	}
	Key = getKey
	return nil
}
