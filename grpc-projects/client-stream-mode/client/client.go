package main

import (
	pb "client-rpc-modes/proto"
	"context"
	"log"
	"strconv"

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
	stream, err := client.SendClientStream(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SendClientStream: %s", err)
	}
	for i := 0; i < 10; i++ {
		// time.Sleep(1 * time.Second)
		if err := stream.Send(&pb.HelloRequest{RequestString: "Hello from the client sent the message " + strconv.Itoa(i+1) + " time"}); err != nil {
			log.Fatalf("Error when sending a message from client: %s", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error when closing and receiving response: %s", err)
	}
	log.Printf("Response from server: %s", res.ResponseString)
}
