package main

import (
	"net/http"
	"log"
)

func main() {
    // Serve static files
    http.Handle("/", http.FileServer(http.Dir("public/index.html")))

    // Start server
    const addr = ":8080"
    if err := http.ListenAndServe(addr, nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}