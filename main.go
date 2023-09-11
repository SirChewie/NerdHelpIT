package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/SirChewie/NerdHelpIT/controllers"
	"github.com/SirChewie/NerdHelpIT/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.ParseTemplate(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/faq", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "support.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/support", controllers.StaticHandler(tpl))

	tpl, err = views.ParseTemplate(filepath.Join("templates", "about.gohtml"))
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/about", controllers.StaticHandler(tpl))

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
