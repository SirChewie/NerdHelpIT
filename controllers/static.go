package controllers

import (
	"net/http"

	"github.com/SirChewie/NerdHelpIT/views"
)

func static_handler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
