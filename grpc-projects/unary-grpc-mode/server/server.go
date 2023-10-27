package main

import (
	"context"
	"log"
	"net"
	pb "unary-rpc-modes/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s server) SendMessage(ctx context.Context, message *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received message body from client: %s", message.RequestString)
	return &pb.HelloResponse{ResponseString: "Hello from the server!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}

}
