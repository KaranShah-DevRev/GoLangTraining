package main

import (
	"log"
	"net"
	pb "server-rpc-modes/proto"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s server) SendServerStream(req *pb.HelloRequest, stream pb.ChatService_SendServerStreamServer) error {
	for i := 0; i < 10; i++ {
		res := &pb.HelloResponse{
			ResponseString: req.RequestString + " " + strconv.Itoa(i),
		}
		stream.Send(res)
	}
	return nil
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
