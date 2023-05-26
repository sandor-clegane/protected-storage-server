package main

import (
	"context"
	"log"

	"protected-storage-server/proto"

	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewGrpcServiceClient(conn)

	_, err = client.CreateUser(context.Background(), &proto.UserRegisterRequest{
		Login:    "test",
		Password: "test",
	})
	if err != nil {
		log.Println(err)
	}
}
