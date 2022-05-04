package client

import (
	"html/template"
	"log"
	"net/http"
)

func serveStatic(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.String())
	t := template.Must(template.ParseGlob("../../spa/*.html"))
	template.Must(t.ParseGlob("../../spa/js/*.js"))
}

func Run() {
	http.HandleFunc("/", serveStatic)
	log.Println("l&s localhost:5000")
	http.ListenAndServe("localhost:5000", nil)
}
