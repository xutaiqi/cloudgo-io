package service

import (
	"net/http"

	"text/template"

	"github.com/unrolled/render"
)

func home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/index.html"))
	t.Execute(w, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form["username"][0]
	password := r.Form["password"][0]

	t := template.Must(template.ParseFiles("templates/info.html"))
	t.Execute(w, map[string]string{
		"username": username,
		"password": password,
	})

}

func homeHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.HTML(w, http.StatusOK, "index", struct {
			ID      string `json:"id"`
			Content string `json:"content"`
		}{ID: "8675309", Content: "Hello from Go!"})
	}
}
