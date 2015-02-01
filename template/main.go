package main

import (
	"html/template"
	"io"
	"os"
)

type Data struct {
	Name string
	Age  int
	Info string
}

var templates = map[string]interface {
	Execute(io.Writer, interface{}) error
}{}

func parseTemplate(sets [][]string) {
	for _, set := range sets {
		var t = template.Must(template.ParseFiles(set...))
		t = t.Lookup("Root")
		templates[set[0]] = t
	}
}

func main() {
	var result = &Data{Name: "Toomore"}
	parseTemplate([][]string{
		{"main.tpl", "layout.tpl"},
	})
	templates["main.tpl"].Execute(os.Stdout, result)
}
