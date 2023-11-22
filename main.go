package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sunnymotiani/lenslocked/controllers"
	"github.com/sunnymotiani/lenslocked/templates"
	"github.com/sunnymotiani/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))
	r.Get("/signup", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {

		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("starting server on :3000...")
	http.ListenAndServe(":3000", r)

}
