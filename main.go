package main

import (
	"html/template"
	"os"
	"time"

	"github.com/go-shiori/dom"
	distiller "github.com/markusmobius/go-domdistiller"
)

type Page struct {
	Title string
	Body  template.HTML
}

func main() {
	url := "https://betterprogramming.pub/a-programmers-regret-neglecting-math-at-university-9d937655752b"

	// Start distiller
	result, err := distiller.ApplyForURL(url, time.Minute, nil)
	if err != nil {
		panic(err)
	}

	rawHTML := dom.OuterHTML(result.Node)

	// parse html template from file assets/template.html
	tmpl, err := template.ParseFiles("./assets/template.html")
	if err != nil {
		panic(err)
	}

	// create page struct
	page := Page{
		Title: "Go-DomDistiller",
		Body:  template.HTML(rawHTML),
	}

	// execute template
	err = tmpl.ExecuteTemplate(os.Stdout, "template.html", page)
	if err != nil {
		panic(err)
	}

}
