package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

type helloWorldResponse struct {
	Message string `json:"message"`
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {

	// Define port at the begin of program
	port := 8080

	http.HandleFunc("/helloworld", log(HelloWorldHandlerFunc))
	http.HandleFunc("/hello", log(validate(HelloHandlerFunc)))
	fmt.Printf("Server starting on port %v\n", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)

	// With default server, we don't have any error return
	/*----------------- This comment for learning -----------------------
	if errServer != nil {
		panic(errServer.Error())
	}
	--------------------------------------------------------------------*/

	n := 0
	r := retrier.New(retrier.ConstantBackoff(3, 1*time.Second), nil)

	err := r.Run(func() error {
		fmt.Println("Attempt: ", n)
		n++
		return fmt.Errorf("Failed")
	})

	if err != nil {
		fmt.Println(err)
	}
}

func HelloWorldHandlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world\n")
}

// func log() is internal function, returns logging for work flow on program.
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

