package main

import (
	"fmt"
	"net/http"
)

func home_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Hello</h1>")
}

func contact_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>Contact us at <a href=\"mailto:jblalock@nerdhelpit.com\">jblalock@nerdhelpit.com</a>.")
}
func faq_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>FAQ Page</h1>")
}
func not_found_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.NotFound(w, r)

}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		home_handler(w, r)
	case "/contact":
		contact_handler(w, r)
	case "/faq":
		faq_handler(w, r)
	default:
		// TODO: handle page not found error
		not_found_handler(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", router)
}
