package main

import (
	pb "bidrect-modes/proto"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s server) SendBiDirectionalStream(stream pb.ChatService_SendBiDirectionalStreamServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("No more requests from the client")
			break
		}
		if err != nil {
			log.Fatalf("Error when receiving request: %s", err)
		}
		log.Printf("Request from client: %s", req.RequestString)

		// Send a response to the client using the Send method
		if err := stream.Send(&pb.HelloResponse{ResponseString: "Hello from the server"}); err != nil {
			log.Fatalf("Error when sending response: %s", err)
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}
	log.Println("Server started on port 9000")
	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}
