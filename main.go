package main

import (
	"fmt"
	"net/http"
)

// helloHandler handles HTTP requests and responds with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	// Use http.ListenAndServe to start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
