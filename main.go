package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/SirChewie/NerdHelpIT/views"
	"github.com/go-chi/chi/v5"
)

func execute_template(w http.ResponseWriter, filepath string) {
	t, err := views.ParseTemplate(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
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

	tpl, err := views.ParseTemplate(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/", controllers.static_handler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}
	r.Get("/contact", controllers.static_handler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/faq", controllers.static_handler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "support.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/support", controllers.static_handler(tpl))
	fmt.Println("Starting server on :3000")

	http.ListenAndServe(":3000", r)
}
