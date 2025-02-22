package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handle incoming requests and write a response to client
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, client!")
}
