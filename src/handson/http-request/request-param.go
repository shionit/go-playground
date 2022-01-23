package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/microcosm-cc/bluemonday"
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
)

func handler(w http.ResponseWriter, r *http.Request) {
	p := bluemonday.UGCPolicy()
	you := r.FormValue("p")
	you = p.Sanitize(you)
	_, err := fmt.Fprintf(w, "%sさんの運勢は、「%s」です！", you, omikuji[rand.Intn(len(omikuji))])
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
