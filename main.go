package main

import (
    "log"
    "net/http"
    "time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    http.HandleFunc("/", helloHandler)

    server := &http.Server{
        Addr:         ":8080",
        Handler:      nil,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Fatal(server.ListenAndServe())
}
