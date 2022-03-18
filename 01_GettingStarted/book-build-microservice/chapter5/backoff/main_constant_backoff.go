package main

import (
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/retrier"
)

func main() {

	n := 0
	// ConstantBackoff generates a simple back-off strategy of retrying 'n' times,
	// and waiting 'amount' time after each one.
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
