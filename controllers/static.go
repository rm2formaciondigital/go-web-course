package controllers

import (
	"net/http"

	"github.com/rm2formaciondigital/go-web-course/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
