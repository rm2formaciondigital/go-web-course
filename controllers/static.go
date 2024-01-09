package controllers

import (
	"html/template"
	"net/http"

	"github.com/rm2formaciondigital/go-web-course/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "there a free versión?",
			Answer:   "Yes!, We Offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   `We hace support staff answering emails 24/7, though response times may be a bit slower on weekends.`,
		},
		{
			Question: "How do I contact support?",
			Answer:   "Email us - support@lenslocked.com",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
