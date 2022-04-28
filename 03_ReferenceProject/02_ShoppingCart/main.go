package main

import "github.com/nicholasjackson/env"

// declare environment server
var bindAddress = env.String("BIND_ADDRESS", false, ":8080", "Bind address for the server")

func main() {

	env.Parse()
}
