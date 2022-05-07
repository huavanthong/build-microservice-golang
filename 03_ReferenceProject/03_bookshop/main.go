package main

import (
	"bookshop/client"
	"bookshop/server"
	"fmt"
)

func main() {

	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformGetBookList(c)

	fmt.Println(reply)
}
