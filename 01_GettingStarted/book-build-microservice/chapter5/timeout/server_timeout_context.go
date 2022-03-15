// Reference: https://stackoverflow.com/questions/58736588/http-server-handlefunc-loop-on-timeout
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

var WriteTimeout = 1 * time.Second

func main() {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: WriteTimeout + 10*time.Millisecond, //10ms Redundant time
		IdleTimeout:  15 * time.Second,
	}
	router.HandleFunc("/", home)
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("responding\n")
	ctx, _ := context.WithTimeout(context.Background(), WriteTimeout)
	worker, cancel := context.WithCancel(context.Background())
	var buffer string
	go func() {
		// do something
		time.Sleep(2 * time.Second)
		buffer = "ready all response\n"
		//do another
		time.Sleep(2 * time.Second)
		cancel()
		fmt.Printf("worker finish\n")
	}()
	select {
	case <-ctx.Done():
		//add more friendly tips
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Timeout\n")
		return
	case <-worker.Done():
		w.Write([]byte(buffer))
		fmt.Fprint(w, "Worker done\n")
		fmt.Printf("writed\n")
		return
	}
}
