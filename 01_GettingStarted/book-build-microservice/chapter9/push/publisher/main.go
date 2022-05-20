package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	nats "github.com/nats-io/nats.go"
)

var natsServer = flag.String("nats", "4223", "NATS server URI")
var natsClient *nats.Conn

type product struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

func init() {
	flag.Parse()
}

func main() {
	var err error
	// Connect to a server
	natsClient, err = nats.Connect("nats://" + *natsServer)
	if err != nil {
		log.Fatal(err)
	}
	defer natsClient.Close()

	// Implement http handler
	http.DefaultServeMux.HandleFunc("/product", productsHandler)

	// List to port
	log.Println("Starting product write service on port 8080")
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}

func productsHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		insertProduct(rw, r)
	}
}

func insertProduct(rw http.ResponseWriter, r *http.Request) {
	log.Println("/insert handler called")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Simple Publisher
	natsClient.Publish("product.inserted", data)
}
