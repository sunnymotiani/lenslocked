package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to SUNNY SUNNY my my my site HAHAHA site!!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>CONTACT PAGE</h1><p>To Get in touch , email me at </p><a href=\"mailto:sun.motiani@gmail.com\">sun.motiani@gmail.com</a>.")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li>Is There a Free Version? Yes Absolutely !!</li>
		<li>Is Sup[port Service Available ? Yes our support staff are available 24*7</li>
		<li>Contact Options ? Email us - <a href="mailto:sun.motiani@gmail.com">sun.motiani@gmail.com</a></li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", r)

}
