package server

import (
	pb "bookshop/proto/pb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// define port
const port = 1234

func main() {
	log.Printf("Server starting on port %v\n", port)
	StartServer()
}

func StartServer() {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
	}
	defer listener.Close()

	gs := grpc.NewServer()

	b := NewBookshop()

	reflection.Register(gs)
	pb.RegisterInventoryServer(gs, b)
	if err := gs.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type Bookshop struct {
	pb.UnimplementedInventoryServer
}

func NewBookshop() *Bookshop {
	return &Bookshop{}
}

func (s *Bookshop) GetBookList(ctx context.Context, in *pb.GetBookListRequest) (*pb.GetBookListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*pb.Book {
	sampleBooks := []*pb.Book{
		{
			Title:     "The Hitchhiker's Guide to the Galaxy",
			Author:    "Douglas Adams",
			PageCount: 42,
		},
		{
			Title:     "The Lord of the Rings",
			Author:    "J.R.R. Tolkien",
			PageCount: 1234,
		},
	}
	return sampleBooks
}
