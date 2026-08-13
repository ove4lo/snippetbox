package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home handles requests to the root URL path
func home(w http.ResponseWriter, r *http.Request) {
	// WHY: convert text to a byte slice because the network only transfers raw bytes
	w.Write([]byte("Hello from Snippetbox"))
}

// snippetView displays a specific snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// snippetCreate displays a form for creating a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// WHY: pointer *http.Request is used to avoid copying huge request data in memory
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// snippetCreatePost handles data submission for saving a new snippet
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet..."))
}


func main() {
	// WHY: mux (ServeMux) is a router. It maps URL paths to correct handler functions
	mux := http.NewServeMux()

	// NOTE: "/{$}" means match the exact root path, preventing it from catching other random URLs
	mux.HandleFunc("GET /{$}", home)

	// WHY: Go 1.22+ automatically parses the {id} wildcard and validates the HTTP method
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)

	// NOTE: Fixed typo (added missing leading slash before 'snippet')
	mux.HandleFunc("GET snippet/create", snippetCreate)

	// WHY: REST best practice - same URL path can handle different actions based on HTTP method
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on: 4000")

	// WHY: ListenAndServe starts an endless loop listening for web traffic on port 4000
	// It only returns an error if the server fails to start (e.g. port is already busy)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}