package service

import (
	"context"
	"log"

	pb "algogrit.com/simple-grpc/api"
)

type greetingServiceV1 struct {
	pb.UnimplementedSimpleServer
}

func (svc *greetingServiceV1) Greeting(ctx context.Context, hr *pb.HelloRequest) (*pb.GreetingResponse, error) {
	log.Printf("Received: %#v\n", *hr)
	return &pb.GreetingResponse{Message: "Hello, " + hr.Name}, nil
}

func NewGreetingServiceV1() pb.SimpleServer {
	return &greetingServiceV1{}
}
