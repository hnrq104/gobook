package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"os"
)

const naoama = "/home/hq/downloads/naoama.jpg"

func getImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		log.Print(err)
		return nil, err
	}

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return img, nil
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	img, err := getImage(naoama)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	err = jpeg.Encode(w, img, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
}
