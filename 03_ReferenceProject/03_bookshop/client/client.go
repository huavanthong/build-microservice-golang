package client

import (
	pb "bookshop/proto/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateClient() *grpc.ClientConn {
	// Dial creates a client connection to the given target
	client, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect %v", err)
	}
	return client
}

func PerformGetBookList(conn *grpc.ClientConn) *pb.GetBookListResponse {

	client := pb.NewInventoryClient(conn)

	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
	if err != nil {
		log.Fatalf("failed to get book list: %v", err)
	}
	log.Printf("book list: %v", bookList)

	return bookList
}

// func main() {
// 	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("failed to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewInventoryClient(conn)
// 	bookList, err := client.GetBookList(context.Background(), &pb.GetBookListRequest{})
// 	if err != nil {
// 		log.Fatalf("failed to get book list: %v", err)
// 	}
// 	log.Printf("book list: %v", bookList)
// }
