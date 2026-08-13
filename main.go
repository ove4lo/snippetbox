package main

import (
	"log"
	"net/http"
)

// home handles requests to the root URL path
func home(w http.ResponseWriter, r *http.Request) {
	// WHY: convert text to a byte slice because the network only transfers raw bytes
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// WHY: mux (ServeMux) is a router. It maps URL paths to correct handler functions
	mux := http.NewServeMux()

	// NOTE: when someone visits the main page "/", trigger the home() function
	mux.HandleFunc("/", home)

	log.Print("starting server on: 4000")

	// WHY: ListenAndServe starts an endless loop listening for web traffic on port 4000
	// It only returns an error if the server fails to start (e.g. port is already busy)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}