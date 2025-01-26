package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// helloHandler handles HTTP requests and responds with a greeting message.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set security headers
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
	w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")

	// Respond with a greeting message
	fmt.Fprintln(w, "Hello, World!")
}

// hiddenFileMiddleware blocks access to hidden files.
func hiddenFileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, ".") {
			http.Error(w, "403 Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Use a custom handler with security headers and hidden file blocking
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	// Wrap the handler with the hidden file middleware
	handler := hiddenFileMiddleware(mux)

	// Create a custom server with timeouts
	server := &http.Server{
		Addr:           ":8080",
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB max header size
	}

	// Start the HTTP server
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
