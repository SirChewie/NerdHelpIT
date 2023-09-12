package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SirChewie/NerdHelpIT/controllers"
	"github.com/SirChewie/NerdHelpIT/templates"
	"github.com/SirChewie/NerdHelpIT/views"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	tpl, err := views.ParseFS(templates.FS, "layout-page.gohtml", "home.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "contact.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "faq.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/faq", controllers.FAQ(tpl))

	tpl, err = views.ParseFS(templates.FS, "support.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/support", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "about.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/about", controllers.StaticHandler(tpl))

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
