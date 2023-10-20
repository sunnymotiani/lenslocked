package main

import (
	"fmt"
	"net/http"
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

/*
	func pathHandler(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			homeHandler(w, r)
		case "/contacts":
			contactHandler(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}
*/

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contacts":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", router)

}

//TODO: ABC ABC
