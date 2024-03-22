package controllers

import (
	"html/template"
	"net/http"

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
