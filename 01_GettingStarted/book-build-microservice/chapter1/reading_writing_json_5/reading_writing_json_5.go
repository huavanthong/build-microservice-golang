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

type helloWorldRequest struct {
	Name string `json:"name"`
}

const port = 8080

func main() {
	server()
}

func server() {
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	// Like Marshall, we can use NewDecoder() func to create object for decoding.
	// Create object decoder to decode Body request.
	decoder := json.NewDecoder(r.Body)
	//  Decode Body request and put to request.
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// After that, we can use it directly.
	response := helloWorldResponse{Message: "Hello " + request.Name}

	// Create object encoder.
	encoder := json.NewEncoder(w)
	// Write JSON directly to response.
	encoder.Encode(response)
}
