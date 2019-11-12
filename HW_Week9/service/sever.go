package service

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

//NewServer new server
func NewServer() *negroni.Negroni {

	formattor := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formattor)
	n.UseHandler(mx)

	return n
}

//initRoutes init routes.
func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/templates", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/test", apiTestHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/unknown", NotImplemented).Methods("GET")
	mx.HandleFunc("/templates", home).Methods("GET")
	mx.HandleFunc("/templates", login).Methods("POST")
	//mx.PathPrefix("/templates").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(webRoot+"/assets/"))))
	//mx.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir(webRoot))))
}
