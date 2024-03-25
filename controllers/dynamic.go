package controllers

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/SirChewie/NerdHelpIT/models"
	"github.com/SirChewie/NerdHelpIT/templates"
	"github.com/SirChewie/NerdHelpIT/views"
)

func ContactFormHandler(tpl views.Template, es *models.EmailService, recevingEmail, sendingEmail string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}
		emptySpace := []byte("")
		// HTML sanitizer.
		sMessage := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("message")), emptySpace))
		sName := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("name")), emptySpace))
		sPhone := string(regexp.MustCompile(`(?i)<[^>]*>`).ReplaceAll([]byte(r.FormValue("phone")), emptySpace))
		// Email form to SMTP_ADDRESS
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

func CookieFormHandler(tpl views.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tpl.Execute(w, nil)
			return
		}
		if r.FormValue("analytics") != "true" {
			setCookieHandler(w, "false")
			emptyTpl, err := template.ParseFS(templates.FS, "bannerClose.gohtml")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = emptyTpl.Execute(w, nil)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		analyticsTpl, err := template.ParseFS(templates.FS, "analytics.gohtml")
		setCookieHandler(w, "true")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = analyticsTpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
func setCookieHandler(w http.ResponseWriter, val string) {
	cookie := http.Cookie{
		Name:  "analyticsStatus",
		Value: val,
		Path:  "/",

		// MaxAge=0 means no "Max-Age" attribute specified.
		// MaxAge<0 means delete cookie now, equivalent to "MaxAge: 0".
		// MaxAge>0 means Max-Age attribute present and given in seconds.
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)
}

func getCookieHandler(w http.ResponseWriter, r *http.Request, cookieName string) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			http.Error(w, "cookie not found", http.StatusBadRequest)
		default:
			log.Println(err)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	}
	// TODO: Take out
	fmt.Println(cookie)
}
