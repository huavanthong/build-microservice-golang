package main

import (
	"fmt"

	"chapter1/rpc/client"
	"chapter1/rpc/server"
)

func main() {
	// Create a routine to start server.
	go server.StartServer()

	// Create multiple client to server instead of a client communicate to server.
	for i := 0; i < 10; i++ {
		// Boom: Ngay tại đây, là cách mà nó connect server nè cha ơi !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
		c := client.CreateClient()
		// Close connect to server
		defer c.Close()

		// Client make a request
		reply := client.PerformRequest(c)

		fmt.Println(reply.Message)
	}
}
