package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "algogrit.com/simple-grpc/api"

	"algogrit.com/simple-grpc/greeting/service"
)

func main() {
	tcpSkt, err := net.Listen("tcp", ":51114")

	if err != nil {
		log.Fatalln("Unable to open socket:", err)
	}

	server := grpc.NewServer()

	greetingService := service.NewGreetingServiceV1()

	pb.RegisterSimpleServer(server, greetingService)

	log.Println("Starting the server...")
	if err := server.Serve(tcpSkt); err != nil {
		log.Fatal("Unable to serve:", err)
	}
}
