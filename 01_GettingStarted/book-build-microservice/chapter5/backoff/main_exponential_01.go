package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

func main() {

	n := 0
	// ExponentialBackoff generates a simple back-off strategy of retrying 'n' times,
	// and doubling the amount of time waited after each one.
	r := retrier.New(retrier.ExponentialBackoff(3, 1*time.Second), nil)

	err := r.Run(func() error {
		fmt.Println("Attempt: ", n)
		n++
		return fmt.Errorf("Failed")
	})

	if err != nil {
		fmt.Println(err)
	}
}
