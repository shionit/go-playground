package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	omikuji = []string{
		"大大吉",
		"大吉",
		"中吉",
		"小吉",
		"末吉",
		"吉",
		"凶",
		"大凶",
	}
	tmpl *template.Template
)

type Result struct {
	Name    string
	Omikuji string
}

func handler(w http.ResponseWriter, r *http.Request) {
	result := &Result{
		Name:    r.FormValue("p"),
		Omikuji: omikuji[rand.Intn(len(omikuji))],
	}
	err := tmpl.Execute(w, result)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	tmpl = template.Must(template.New("main").
		Parse("<html><body>{{.Name}}さんの運勢は、「<b>{{.Omikuji}}</b>」です！</body></html>"))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
