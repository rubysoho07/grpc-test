package main

import (
	"context"
	"log"
	"net"

	pb "yg_test_service"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedYGTestServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.YGRequest) (*pb.YGResponse, error) {
	log.Printf("Received: %v", in.GetReqMessage())
	return &pb.YGResponse{ResMessage: "Hello " + in.GetReqMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterYGTestServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
