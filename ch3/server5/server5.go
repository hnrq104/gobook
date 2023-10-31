package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	var color string
	if r.Form.Has("c") {
		color = r.Form["c"][0]
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	ProduceSVG(w, color)
	// fmt.Fprintf(w, "%s\n", color)
}
