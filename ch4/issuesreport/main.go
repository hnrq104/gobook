package main

import (
	"gobook/ch4/github"
	"log"
	"os"
	"text/template"
	"time"
)

const templ1 = `{{.TotalCount}} issues :
{{range .Items}}---------------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ1))

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("issuesreport: %v", err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatalf("issuesreport: %v", err)
	}
}
