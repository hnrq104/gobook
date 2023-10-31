package main

//issues server serves the html found in issues html, cacheing the result

import (
	"fmt"
	"gobook/ch4/github"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// will stay at localhost:8000/
var issueList = template.Must(template.New("issuelist").Parse(`
<h1>{{.TotalCount}} issues </h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
  <th>server</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'> {{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td><a href='localhost:8000/?number={{.Number}}'>link</a></td>
</tr>
{{end}}
</table>
`))

// will stay at localhost:8000/number
var issueReport = template.Must(template.New("issueTempo").Parse(`
<h1>{{.Title}} #{{.Number}} </h1>
<h2><a href='{{.User.HTMLURL}}'> User: {{.User.Login}} </a></h2>
<p>This issue can be found <a href='{{.HTMLURL}}'>{{.HTMLURL}}</a>!</p>
<h2>Issue</h2>
<p>{{.Body}}</p>
<p>Created at {{.CreatedAt}}</p>
`))

type Cache struct {
	github.IssueSearchResult
	ind map[string]*github.Issue
}

func (c Cache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if nothing, just print list
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Problem parsing form")
		log.Print("Problem parsing form")
	} else if r.Form.Has("number") {
		issue, ok := c.ind[strings.Join(r.Form["number"], "")]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "ERROR 404: Couldn't find %q", r.Form["number"])
			log.Print("Could not find issue")
		} else {
			if err := issueReport.Execute(w, issue); err != nil {
				log.Printf("%v\n", err)
			}
		}
	} else if err := issueList.Execute(w, c.IssueSearchResult); err != nil {
		log.Print("Problem executing html")
	}
}

func main() {
	fmt.Printf("Starting web server at localhost:8000\n")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	var c Cache
	c.IssueSearchResult = *result
	c.ind = make(map[string]*github.Issue)

	//adds each issue to map
	for _, i := range c.Items {
		c.ind[strconv.Itoa(i.Number)] = i
	}

	http.Handle("/", c)
	log.Fatal(http.ListenAndServe("localhost:8000", c))
}
