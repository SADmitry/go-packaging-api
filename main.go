package main

import (
	"fmt"
	"net/http"
)

// main initializes the application, loads configuration, and starts the HTTP server.
func main() {
	loadConfig()
	http.HandleFunc("/calculate", handler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
