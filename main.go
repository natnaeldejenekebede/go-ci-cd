package main

import (
	"fmt"
	"net/http"
	"time"
)

// helloHandler handles HTTP requests and responds with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	// Create a custom server with timeouts to avoid Gosec warnings
	server := &http.Server{
		Addr:           ":8080",
		Handler:        nil, // Default to use http.HandleFunc
		ReadTimeout:    10 * time.Second, // Read timeout for incoming requests
		WriteTimeout:   10 * time.Second, // Write timeout for responses
		MaxHeaderBytes: 1 << 20, // 1MB max header size
	}

	// Start the HTTP server with the custom timeouts
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
