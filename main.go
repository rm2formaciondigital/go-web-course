package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a "+"href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.")
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h2>Q: Is there a free versión</h2>
	<h2>A: Yes: We Offer a free trial for 30 days on any paid plans.</h2>`)
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
