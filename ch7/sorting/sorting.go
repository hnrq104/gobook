package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
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

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sorting order
type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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

func main() {
	buf := bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to sorter! specify by what column you want it sorted")
	for buf.Scan() {
		switch strings.ToLower(buf.Text()) {
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
		printTracks(tracks)
		fmt.Println()
		fmt.Print("Next column: ")
	}
}

//For execise 7.8 i would implement a queue and keep track of who is first.

//listrack example

// func listTrack(db *sql.DB, artist string, minYear, maxYear int) {
// 	result, err := db.Exec(
// 		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
// 		artist,minYear,maxYear)
// }
