package main

import (
	"fmt"
	"net/http"
)

const port = ":3000"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Starting server on %s\n", port)
	http.ListenAndServe(port, nil)
}
