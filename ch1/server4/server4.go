package main

import (
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		return
	}

	p := make(map[string]int)

	for k, v := range r.Form {
		if number, err := strconv.Atoi(v[0]); err == nil {
			p[k] = number
		}

	}

	parametrized_lj(w, p)

}
