package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/deadline"
)

func main() {

	port := 8080

	http.HandleFunc("/slow", makeNormalRequest)
	http.HandleFunc("/timeout", makeTimeoutRequest)

	fmt.Printf("Server starting on port %v\n", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

}

func makeNormalRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "slow Funtion is starting\n")
	slowFunction()
}

func makeTimeoutRequest(w http.ResponseWriter, r *http.Request) {

	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()
		return nil
	})

	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
		fmt.Fprint(w, "Timeout\n")
	default:
		fmt.Println(err)
		fmt.Fprint(w, err)
	}
}

func slowFunction() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(1 * time.Second)
	}
}
