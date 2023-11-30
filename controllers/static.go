package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/SirChewie/NerdHelpIT/models"
	"github.com/SirChewie/NerdHelpIT/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func GalleryHandler(tpl views.Template) http.HandlerFunc {
	sites := []struct {
		// !Warning! only user template.HTML if you trust the input. This is a security risk. If you don't trust the input, use "string" instead.
		Site        template.HTML
		Description string
	}{
		{
			Site:        "<a class=\"flex justify-center items-center\" href=\"https://pazitivo.com\" target=\"_blank\"><img class=\"rounded-t-lg\" alt=\"Customer's webpage.\" src=\"assets\\PazitivoLogo.jpeg\"/></a>",
			Description: "Pazitivo Inspired Kitchen",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, sites)
	}

}

func ContactFormHandler(tpl views.Template, es *models.EmailService, recevingEmail, sendingEmail string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}
		// TODO: Sanatize the form data

		email := models.Email{

			To:        recevingEmail,
			From:      sendingEmail,
			Subject:   "New form submission from: " + r.FormValue("subject"),
			PlainText: "Email: " + r.FormValue("email") + "\n" + r.FormValue("message"),
			HTML:      `<span>Email:` + r.FormValue("email") + `<br>` + r.FormValue("message") + `</span>`,
		}
		err := es.Send(email)
		if err != nil {
			fmt.Printf("error sending email: %v", err)
			return
		}
		tpl.Execute(w, struct{ Success bool }{true})
		return
	}
}
