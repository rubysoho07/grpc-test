package main

import (
	"context"
	"log"
	"os"
	"time"
	pb "yg_test_service"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("not connected: %v", err)
	}
	defer conn.Close()
	c := pb.NewYGTestServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.YGRequest{ReqMessage: name})
	if err != nil {
		log.Fatalf("Could not send message: %v", err)
	}

	log.Printf("Greeting: %s", r.GetResMessage())
}
