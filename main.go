package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rm2formaciondigital/go-web-course/controllers"
	"github.com/rm2formaciondigital/go-web-course/templates"
	"github.com/rm2formaciondigital/go-web-course/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gotmpl"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gotmpl"))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "faq.gotmpl"))))

	r.Get("/support", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "support.gotmpl"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
