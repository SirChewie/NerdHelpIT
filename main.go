package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func home_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
func contact_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>Contact us at <a href=\"mailto:jblalock@nerdhelpit.com\">jblalock@nerdhelpit.com</a>.")
}
func faq_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>FAQ Page</h1>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", home_handler)
	r.Get("/contact", contact_handler)
	r.Get("/faq", faq_handler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
