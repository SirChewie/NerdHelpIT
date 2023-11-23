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
	esUsername := os.Getenv("SMTP_USERNAME")
	esPassword := os.Getenv("SMTP_PASSWORD")

	// Setup the email service.
	models.NewEmailService(models.SMTPConfig{
		Host:     esHost,
		Port:     esPort,
		Username: esUsername,
		Password: esPassword,
	})

	/////////////////////////////////////////
	// This is just for testing purposes only, set the above NewEmailService to "es :=" for the test to work.
	//err = es.EmailForm(os.Getenv("SMTP_SENTADDRESS"), username)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Email Sent successfully")
	////////////////////////////////////////

	// Declares the router interface
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
	r.Get("/contact", controllers.Contact(tpl))

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

	fmt.Println("Starting server on:", sAddress)
	err = http.ListenAndServe(sAddress, r)
	if err != nil {
		log.Fatalf("Error Serving:%s\nError:%v", sAddress, err)
	}
}
