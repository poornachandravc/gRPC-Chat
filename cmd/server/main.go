package main

import (
	"context"
	"fmt"
	pb "grpc-chat/api/v1"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedMessageServer
}

func (s *server) SendMessage(ctx context.Context, input *pb.RequestMessage) (*pb.ResponseMessage, error) {
	log.Printf("Received %s", input.GetName())
	return &pb.ResponseMessage{Message: "Hello " + input.GetName()}, nil
}

func SayHello() string {
	return "started server"
}

func main() {
	message := SayHello()
	fmt.Println(message)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
