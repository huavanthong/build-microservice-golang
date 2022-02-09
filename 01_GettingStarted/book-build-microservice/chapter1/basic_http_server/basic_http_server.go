package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8080

	// The HandleFunc method creates a Handler type on the DefaultServeMux handler,
	// mapping the path passed in the first parameter to the function in the second parameter.
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)

	// Function: ListenAndServe takes two parameters, the TCP network address to bind the server to and the handler that will be used to route requests
	// @Para1: port -- network address 8080 bind the server to all available IP addresses on port 8080.
	// @Para2: nil  -- this is because we are using the DefaultServeMux handler, which we are setting up with our call to http.
	// Since ListenAndServe blocks if the server starts correctly we will never exit on a successful start.
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
