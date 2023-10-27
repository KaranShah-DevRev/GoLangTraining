package main

import (
	pb "client-rpc-modes/proto"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s server) SendClientStream(stream pb.ChatService_SendClientStreamServer) error {
	total := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{
				ResponseString: "Total recieved messages: " + strconv.Itoa(total),
			})
		}
		if err != nil {
			return err
		}
		total++
		fmt.Printf("Received from client the message: %v %d times\n", req.RequestString, total)
	}
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
