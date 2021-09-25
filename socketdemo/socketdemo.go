package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	println("Starting Server")
	println("Starting Websocket Server")
	go manager.start()

	http.HandleFunc("/", handler)
	http.HandleFunc("/ws", wsPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *PageData) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", &PageData{Title: "Teste"})
}
