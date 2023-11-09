package controllers

import (
	"html/template"
	"net/http"

	"github.com/sunnymotiani/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	question := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free Version ?",
			Answer:   "Yes We do offer a 30 Day Free trial",
		},
		{
			Question: "What Are you support hours ?",
			Answer:   "24/7",
		},
		{
			Question: "Contact Options ?",
			Answer:   ` Email us - <a href="mailto:sun.motiani@gmail.com">sun.motiani@gmail.com</a>`,
		},
		{
			Question: "Office Location?",
			Answer:   "our entire team is remote",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, question)
	}
}
