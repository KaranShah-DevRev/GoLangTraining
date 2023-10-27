package main

import (
	pb "bidrect-modes/proto"
	"context"
	"io"
	"log"
	"time"

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
	stream, err := client.SendBiDirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SendBidirectionalStream: %s", err)
	}
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.HelloRequest{RequestString: "Hello from the client sent the message "}); err != nil {
			log.Fatalf("Error when sending a message from client: %s", err)
		}
	}
	println("Closing client stream")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("No more responses from server")
			break
		}
		if err != nil {
			log.Fatalf("Error when receiving response: %s", err)
		}
		log.Printf("Response from server: %s", req.ResponseString)
	}
}
