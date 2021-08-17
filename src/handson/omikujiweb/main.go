package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	omikuji := []string{
		"大大吉",
		"大吉",
		"中吉",
		"小吉",
		"末吉",
		"吉",
		"凶",
		"大凶",
	}
	fmt.Fprint(w, "あなたの運勢は、"+omikuji[rand.Intn(len(omikuji))])
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
