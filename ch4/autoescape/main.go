package main

import (
	"html/template"
	"log"
	"os"
)

func main() {

	const templ1 = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("autoescape").Parse(templ1))
	var data struct {
		A string        //untrusted plain data
		B template.HTML //trusted html
	}
	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}
