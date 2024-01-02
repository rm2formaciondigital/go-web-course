package main

import (
	"fmt"
	"gocourse/views"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
		return
	}

	viewTpl := views.Template{
		HTMLTpl: tpl,
	}

	viewTpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gotmpl")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gotmpl")
	executeTemplate(w, tplPath)
}

func contactIdHandler(w http.ResponseWriter, r *http.Request) {
	contactID := chi.URLParam(r, "contactID")

	ctx := r.Context()
	key, ok := ctx.Value("key").(string)
	if !ok {
		key = "clave no definida"
	}

	w.Write([]byte(fmt.Sprintf("Hi %v, %v", contactID, key)))
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gotmpl"))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/contact/{contactID}", contactIdHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
