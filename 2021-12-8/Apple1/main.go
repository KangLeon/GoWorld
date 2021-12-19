package main

import (
	trippb "Apple1/proto/gen/go"
	trip "Apple1/tripservice"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGetway()
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

func startGRPCGetway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux()
	err := trippb.RegisterTripServiceHandlerFromEndpoint(
		c,
		mux,
		":8081",
		[]grpc.DialOption{grpc.WithInsecure()},
	)

	if err != nil {
		log.Fatalf("Cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Cannot listen and server: %v", err)
	}
}
