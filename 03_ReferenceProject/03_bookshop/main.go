package main

import (
	"log"
	"net"

	pb "bookshop/server/pb"

	"bookshop/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	b := server.NewBookshop()

	reflection.Register(s)
	pb.RegisterInventoryServer(s, b)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
