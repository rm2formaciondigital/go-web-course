package controllers

import (
	"net/http"

	"github.com/rm2formaciondigital/go-web-course/views"
)

type Users struct {
	Templates struct {
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
