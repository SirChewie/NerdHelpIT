package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello</h1>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}
