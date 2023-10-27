package main

import (
	"log"

	pb "unary-rpc-modes/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.ChatServiceClient

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()
	client = pb.NewChatServiceClient(conn)
	r := gin.Default()
	r.GET("/sendMessage/:message", clientConnection)
	r.Run(":8000")

}

func clientConnection(c *gin.Context) {
	message := c.Param("message")
	resp, err := client.SendMessage(c, &pb.HelloRequest{RequestString: message})
	if err != nil {
		log.Fatalf("Error when calling SendMessage: %s", err)
	}
	c.JSON(200, gin.H{
		"message": resp.ResponseString,
	})
}
