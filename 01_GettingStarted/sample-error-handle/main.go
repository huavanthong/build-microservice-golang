package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/you/somepkg/handler"
)

func main() {
	db, err := sql.Open("connectionstringhere")
	if err != nil {
		log.Fatal(err)
	}

	// Initialise our app-wide environment with the services/info we need.
	env := &handler.Env{
		DB:   db,
		Port: os.Getenv("PORT"),
		Host: os.Getenv("HOST"),
		// We might also have a custom log.Logger, our
		// template instance, and a config struct as fields
		// in our Env struct.
	}

	// Note that we're using http.Handle, not http.HandleFunc. The
	// latter only accepts the http.HandlerFunc type, which is not
	// what we have here.
	http.Handle("/", handler.Handler{env, handler.GetIndex})

	// Logs the error if ListenAndServe fails.
	log.Fatal(http.ListenAndServe(":8000", nil))
}
