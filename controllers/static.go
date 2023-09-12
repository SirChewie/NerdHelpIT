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

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		// !Warning! only user template.HTML if you trust the input. This is a security risk. If you don't trust the input, use "string" instead.
		Answer template.HTML
	}{
		{
			Question: "What is NerdHelpIT?",
			Answer:   "NerdHelpIT is a web application that helps you find the right help for your problem.",
		},
		{

			Question: "What Services do we offer?",
			Answer:   "NerdHelpIT is a web application that helps you find the right help for your problem.",
		},
		{

			Question: "How do I contact NerdHelpIT?",
			Answer:   "Contact us at <a href=\"mailto:jblalock@nerdhelpit.com\">jblalock@nerdhelpit.com</a> if you have any questions.",
		},
		{

			Question: "Where is your data stored?",
			Answer:   "NerdHelpIT stores all of your data encrypted on our servers within the United States.",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
