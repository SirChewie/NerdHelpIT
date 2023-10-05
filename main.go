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

	// Render the layout page then render the content page.
	tpl, err := views.ParseFS(templates.FS, "tailwind.gohtml", "home.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "contact.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "faq.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/faq", controllers.FAQ(tpl))

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "support.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/support", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "about.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/about", controllers.StaticHandler(tpl))

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
