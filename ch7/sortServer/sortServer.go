package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Now and Then", "The Beatles", "Now and Then", 2023, length("4m6s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Something", "The Beatles", "Abbey Road", 1969, length("3m2s")},
	{"Brain Damage", "Pink Floyd", "The Darkside of the Moon", 1973, length("3m46s")},
	{"Children of the Grave", "Black Sabbath", "Masters of Reality", 1971, length("5m15s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// Using sort stable
//using customSort, sort by artist

func lessArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}

func lessTitle(x, y *Track) bool {
	return x.Title < y.Title
}

func lessAlbum(x, y *Track) bool {
	return x.Album < y.Album
}

func lessYear(x, y *Track) bool {
	return x.Year < y.Year
}

func lessLength(x, y *Track) bool {
	return x.Length < y.Length
}

var musicList = template.Must(template.New("musicList").Parse(`
<table>
<tr style='text-align: left'>
  <th><a href=/title>Title</a></th>
  <th><a href=/artist>Artist</th>
  <th><a href=/album>Album</th>
  <th><a href=/year>Year</th>
  <th><a href=/length>Length</th>
</tr>
{{range .}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.EscapedPath()[1:])
	switch strings.ToLower(r.URL.EscapedPath()[1:]) {
	case "title":
		sort.Stable(customSort{tracks, lessTitle})
	case "artist":
		sort.Stable(customSort{tracks, lessArtist})
	case "album":
		sort.Stable(customSort{tracks, lessAlbum})
	case "year":
		sort.Stable(customSort{tracks, lessYear})
	case "length":
		sort.Stable(customSort{tracks, lessLength})
	}
	musicList.Execute(w, tracks)
}
