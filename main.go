package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a "+"href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h2>Q: Is there a free versión</h2>
	<h2>A: Yes: We Offer a free trial for 30 days on any paid plans.</h2>`)
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "<h1>Page Not Found</h1>")
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on http://localhost:3000")
	http.ListenAndServe(":3000", router)
	// http.HandleFunc("/", homeHandler)
	// http.Handle("/contact", http.HandlerFunc(contactHandler))
	// fmt.Println("Starting the server on http://localhost:3000")
	// http.ListenAndServe(":3000", nil)

	// option 1 - Convert Our handler func in a type Handler
	// var router http.HandlerFunc = pathHandler
	// fmt.Println("Starting the server on :3000...")
	// http.ListenAndServe(":3000", router)

	// http.Handler - Interface with the ServeHTTP method
	//http.HandlerFunc - Function type that accepts same args as ServeHTTP method. Also implements http.Handler

	// Option 2 - Conver func y a type HandlerFunc
	// fmt.Println("Starting the server on :3000...")
	// http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
