package main

import (
	"fmt"
	"log"
	"net/http"
)

// formHandler handles the user request and the server response.
// It processes the form data submitted by the user.
func formHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	// Handle the post request
	fmt.Fprintf(w, "Post request successful\n")

	// Extract data from the form
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Output the extracted data
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// helloHandler responds to requests with a simple greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "METHOD NOT SUPPORTED", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "HELLO!")
}

// logRequest is a middleware that logs the HTTP request paths.
func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested URL: %s\n", r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	// Set up the static file server, serving files out of the './static' directory.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", logRequest(fileServer)) // Wrap the file server with the logging function

	// Register handler functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080 and log any errors
	fmt.Println("Starting server at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
