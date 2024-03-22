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

	// Load configuration from ENV variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:")
	}

	sAddress := os.Getenv("SERVING_ADDRESS")
	esHost := os.Getenv("SMTP_HOST")
	esPortStr := os.Getenv("SMTP_PORT")
	// convert needed variables to int values
	esPort, err := strconv.Atoi(esPortStr)
	if err != nil {
		log.Fatalf("Error parsing port number.\n Provided number: %v\n Error:%v", esPortStr, err)
	}
	esDefaultDropOff := os.Getenv("SMTP_DEFAULT_DROP_OFF")
	esUsername := os.Getenv("SMTP_USERNAME")
	esPassword := os.Getenv("SMTP_PASSWORD")

	// Setup the email service config.
	es := models.NewEmailService(models.SMTPConfig{
		Host:     esHost,
		Port:     esPort,
		Username: esUsername,
		Password: esPassword,
	})

	// Declares the router interface
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
	r.Post("/contact", controllers.ContactFormHandler(tpl, es, esDefaultDropOff, esUsername))

	tpl, err = views.ParseFS(templates.FS, "tailwind.gohtml", "gallery.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}
	r.Get("/gallery", controllers.GalleryHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "cookieSettings.gohtml", "analytics.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}
	r.Get("/cookieSettings", controllers.StaticHandler(tpl))
	r.Post("/cookieSettings", controllers.CookieFormHandler(tpl))

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

	tpl, err = views.ParseFS(templates.FS, "analytics.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}
	r.Get("/analytics", controllers.StaticHandler(tpl))

	tpl, err = views.ParseFS(templates.FS, "empty.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		return
	}
	r.Get("/empty", controllers.StaticHandler(tpl))

	assetsHandler := http.FileServer(http.Dir("assets"))
	r.Get("/assets/*", http.StripPrefix("/assets", assetsHandler).ServeHTTP)

	fmt.Println("Starting server on:", sAddress)
	err = http.ListenAndServe(sAddress, r)
	if err != nil {
		log.Fatalf("Error Serving:%s\nError:%v", sAddress, err)
	}
}
