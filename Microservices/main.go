package main

import (
	h "Microservices/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", h.Home)
	http.HandleFunc("/about", h.About)
	fmt.Printf("Starting application on port %s", portNumber)
	http.ListenAndServe(portNumber, nil)
}
