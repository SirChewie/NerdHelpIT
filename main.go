package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/SirChewie/NerdHelpIT/controllers"
	"github.com/SirChewie/NerdHelpIT/models"
	"github.com/SirChewie/NerdHelpIT/templates"
	"github.com/SirChewie/NerdHelpIT/views"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:")
	}
	host := os.Getenv("SMTP_HOST")
	portStr := os.Getenv("SMTP_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	// This is just for testing purposes only
	err = es.EmailForm("jhb1085@gmail.com", username)
	if err != nil {
		panic(err)
	}
	fmt.Println("Email Sent successfully")

	r := chi.NewRouter()

	// Render the layout page then render the content page.
	tpl, err := views.ParseFS(templates.FS, "home.gohtml")
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

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "gallery.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/gallery", controllers.Gallery(tpl))

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

	tpl, err = views.ParseFS(templates.FS, "nav.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/nav", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "navClose.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}

	r.Get("/navClose", controllers.StaticHandler(tpl))

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", r)
}
