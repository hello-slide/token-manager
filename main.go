package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/hello-slide/token-manager/token"
)

var Key []byte = []byte(os.Getenv("PUBLIC_KEY"))

func createHandler(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	if in.ContentType != "text/plain" {
		return nil, fmt.Errorf("The content-type must be `text/plain.`")
	}

	resultToken, err := token.Create(string(in.Data), Key)
	if err != nil {
		return nil, err
	}

	return &common.Content{
		Data:        []byte(resultToken),
		ContentType: "text/plain",
	}, nil
}

func verifyHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in.ContentType != "text/plain" {
		return nil, fmt.Errorf("The content-type must be `text/plain.`")
	}

	resultToken, err := token.Verify(string(in.Data), Key)
	if err != nil {
		return nil, err
	}

	return &common.Content{
		Data:        []byte(resultToken),
		ContentType: "text/plain",
	}, nil
}

func main() {
	s, err := daprd.NewService(":3000")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	if err := s.AddServiceInvocationHandler("create", createHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.AddServiceInvocationHandler("verify", verifyHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
