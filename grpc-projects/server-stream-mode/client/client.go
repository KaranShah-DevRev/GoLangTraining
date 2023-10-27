package main

import (
	"context"
	"log"
	pb "server-rpc-modes/proto"

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
	stream, err := client.SendServerStream(context.Background(), &pb.HelloRequest{RequestString: "Hello from the client"})
	if err != nil {
		log.Fatalf("Error when calling SendServerStream: %s", err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error when receiving response: %s", err)
		}
		log.Printf("Response from server: %s", res.ResponseString)
	}
}
