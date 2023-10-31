/*
Exercise 4.13: The JSON-based web service of the Open Movie Database lets you search
https://omdbapi.com/ for a movie by name and download its poster image. Write a tool
poster that downloads the poster image for the movie named on the command line.
*/

/*My OMDb APi key : 7323eadc*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const dataUrl = "http://www.omdbapi.com/?apikey=7323eadc&t="

/*
type Critic struct {
	Source string
	Value  string
}

type Movie struct {
	Title    string
	Year     string
	Rated    string
	Released string
	Poster   string
	Actors   []string
	Ratings  []Critic
	Plot     string
	Genre    []string
}
*/

type MoviePoster struct {
	Poster string
}

const ImagePath = "/home/hq/pics"

func main() {
	movieName := strings.Join(os.Args[1:], " ")
	query := url.QueryEscape(movieName)

	resp, err := http.Get(dataUrl + query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}

	var posterURL MoviePoster
	err = json.NewDecoder(resp.Body).Decode(&posterURL)
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "error decoding %s\n", resp.Request.URL)
	}
	resp.Body.Close()

	imgResp, err := http.Get(posterURL.Poster)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get: %v\n", err)
		os.Exit(1)
	}

	filePattern := strings.Join(os.Args[1:], "") + "*"
	file, err := os.CreateTemp(ImagePath, filePattern)
	if err != nil {
		imgResp.Body.Close()
		fmt.Fprintf(os.Stderr, "poster: %v\n", err)
		os.Exit(1)
	}

	_, err = io.Copy(file, imgResp.Body)
	if err != nil {
		imgResp.Body.Close()
		file.Close()
		fmt.Fprintf(os.Stderr, "io.Copy: %v\n", err)
		os.Exit(1)
	}

	imgResp.Body.Close()
	file.Close()
	fmt.Printf("Movie poster downloaded to %s\n", ImagePath)
}
