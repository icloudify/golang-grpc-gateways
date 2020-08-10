package main

import (
	"context"
	"fmt"
	pb "github.com/ravindra031/golang-grpc-gateways/gatewaypb/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":10000"
)

type server struct{}

func (s *server) SendGet(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	fmt.Println("Received GET method " + in.Name)
	return &pb.TemplateResponse{Message: "Received GET method " + in.Name}, nil
}

func (s *server) SendPost(ctx context.Context, in *pb.TemplateRequest) (*pb.TemplateResponse, error) {
	fmt.Println("Received POST method " + in.Name)
	return &pb.TemplateResponse{Message: "Received POST method " + in.Name}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
