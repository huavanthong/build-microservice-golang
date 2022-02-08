package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// type helloWorldResponse struct {
// 	Message string `json:"message"`
// }

// Some feature for usage using json in Golang
type helloWorldResponse struct {
	// change the output field to be "message"
	Message string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output the field if the value is empty
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/helloworld2", helloWorldHandlerAdvantage)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// Normal case: Send one message
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {

	response := helloWorldResponse{Message: "HelloWorld"}

	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

// Advantage case: Send multiple message.
/*
Prolem:
	Đặt vấn đề rằng, nếu chúng ta gửi mutiple json message vào một response. Thì điều gì sẽ xảy ra?
	Và làm cách nào ta có thể handle được multiple message đó.
Output:

Solution

*/
func helloWorldHandlerAdvantage(w http.ResponseWriter, r *http.Request) {

	response1 := helloWorldResponse{Message: "HelloWorld"}

	response2 := helloWorldResponse{Message: "Hello World", Id: 333}

	data1, err := json.Marshal(response1)
	data2, err := json.Marshal(response2)

	if err != nil {
		panic("Ooops")
	}

	fmt.Fprintf(w, string(data1), string(data2))
}
