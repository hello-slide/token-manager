package main

import (
	"context"
	"fmt"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/hello-slide/token-manager/token"
)

func createHandler(ctx context.Context, in *common.InvocationEvent) (*common.Content, error) {
	if in.ContentType != "text/plain" {
		return nil, fmt.Errorf("The content-type must be `text/plain.`")
	}

	resultToken, err := token.Create(string(in.Data))
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

	resultToken, err := token.Verify(string(in.Data))
	if err != nil {
		return nil, err
	}

	return &common.Content{
		Data:        []byte(resultToken),
		ContentType: "text/plain",
	}, nil
}

func init() {
	client, err := dapr.NewClient()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	token.GetKey(&client, &ctx)
}

func main() {
	s, err := daprd.NewService(":50001")
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
