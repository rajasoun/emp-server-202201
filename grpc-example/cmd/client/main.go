package main

import (
	"context"
	"log"
	"time"

	pb "algogrit.com/simple-grpc/api"

	"google.golang.org/grpc"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	conn, err := grpc.DialContext(ctx, ":51114", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Unable to connect to server:", err)
	}

	simple := pb.NewSimpleClient(conn)

	ctx, _ = context.WithCancel(context.Background())

	hr := pb.HelloReq{ID: -100}

	resp, err := simple.Greeting(ctx, &hr)

	if err != nil {
		log.Fatalln("Unable to get greeting:", err)
	}

	log.Println("Received:", resp.Message)
}
