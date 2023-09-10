package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func execute_template(w http.ResponseWriter, filepath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
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

func home_handler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	execute_template(w, tplPath)
}

func contact_handler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	execute_template(w, tplPath)
}

func faq_handler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	execute_template(w, tplPath)
}

func support_handler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "support.gohtml")
	execute_template(w, tplPath)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", home_handler)
	r.Get("/contact", contact_handler)
	r.Get("/faq", faq_handler)
	r.Get("/support", support_handler)
	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
