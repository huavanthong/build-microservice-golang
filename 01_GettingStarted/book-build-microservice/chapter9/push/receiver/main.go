package main

import (
	"encoding/json"
	"flag"
	"log"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

type product struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

var natsClient *nats.Conn

var natsServer = flag.String("nats", "4223", "NATS server URI")

func init() {
	flag.Parse()

}

func main() {
	var err error
	// Connect to a server
	natsClient, err = nats.Connect("nats://" + *natsServer)
	if err != nil {
		log.Println("Connected to " + *natsServer)
		log.Fatal(err)
	}
	defer natsClient.Close()

	// Simple Async Subscriber
	log.Println("Subscribing to events")
	natsClient.Subscribe("product.inserted", handleMessage)

	// Keep the connection alive
	runtime.Goexit()
}

func handleMessage(m *nats.Msg) {
	p := product{}
	err := json.Unmarshal(m.Data, &p)
	if err != nil {
		log.Println("Unable to unmarshal event object")
		return
	}

	log.Printf("Received message: %v, %#v", m.Subject, p)
}
