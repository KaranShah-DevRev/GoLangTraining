package main

import (
	"context"
	"log"
	"time"
	pb "unary-rpc-modes/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	client := pb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.SendMessage(ctx, &pb.HelloRequest{RequestString: "Hello from the client!"})
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}
	log.Printf("Response from server: %s", resp.ResponseString)
}
