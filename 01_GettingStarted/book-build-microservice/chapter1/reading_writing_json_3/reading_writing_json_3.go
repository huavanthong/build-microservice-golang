package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// The ResponseWriter is an interface that defines three methods:
//
// [ ---- Header() ----]
// Returns the map of headers which will be sent by the
// WriteHeader method.
//
// [ ---- Write([]byte) (int, error) ----]
// Writes the data to the connection. If WriteHeader has not
// already been called then Write will call
// WriteHeader(http.StatusOK).
//
// [ ---- WriteHeader(int) ----]
// Sends an HTTP response header with the status code.
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Create message response following member in struct.
	response := helloWorldResponse{Message: "HelloWorld"}
	// Create object encoder.
	encoder := json.NewEncoder(w)
	// Write JSON straight to an open writer
	encoder.Encode(response)
}
