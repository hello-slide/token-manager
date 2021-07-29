package main

import (
	"context"
	"log"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
	"github.com/hello-slide/token-manager/token"
)

func createHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	log.Printf("echo - ContentType:%s, Verb:%s, QueryString:%s, %+v", in.ContentType, in.Verb, in.QueryString, string(in.Data))
	// do something with the invocation here
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}

func verifyHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %v", e.PubsubName, e.Topic, e.ID, e.Data)
	// do something with the event
	return true, nil
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
	s, err := daprd.NewService(":3000")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	if err := s.AddServiceInvocationHandler("create", createHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      "topic1",
	}
	if err := s.AddTopicEventHandler(sub, verifyHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
