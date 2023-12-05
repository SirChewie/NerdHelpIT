package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

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
		emptySpace := []byte("")
		// HTML sanitizer. TODO: iterate over this instead.
		sMessage := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("message")), emptySpace))
		sName := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("name")), emptySpace))
		sPhone := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("phone")), emptySpace))

		email := models.Email{

			To:        recevingEmail,
			From:      sendingEmail,
			Subject:   sName,
			PlainText: "Name: " + sName + "\n" + "Email: " + r.FormValue("email") + "\n" + "Phone number: " + sPhone + "\n" + "Message: " + sMessage + "\n",
			HTML:      `<span>Name: ` + sName + `<br />` + `Email: ` + r.FormValue("email") + `<br />` + `Phone number: ` + sPhone + `<br />` + `Message: ` + `<br />` + sMessage + `<br />` + `</span>`,
		}
		err := es.Send(email)
		if err != nil {
			fmt.Printf("error sending email: %v", err)
			return
		}
		tpl.Execute(w, struct{ Success bool }{true})
	}
}
