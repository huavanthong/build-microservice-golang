package server

import (
	pb "bookshop/server/pb"
	"context"
	"log"
)

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
