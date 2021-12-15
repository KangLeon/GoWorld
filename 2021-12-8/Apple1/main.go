package main

import (
	trippb "Apple1/proto/gen/go"
	trip "Apple1/tripservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listn: %v", err)
	}
	s := grpc.NewServer()
	trippb.RegisterTripServiceServer(
		s,
		&trip.Service{},
	)
	log.Fatal(s.Serve(lis))
}
