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
