package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
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

func serveHome(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req)
	var result = &Data{Name: req.URL.Path}
	parseTemplate([][]string{
		{"main.tpl", "layout.tpl"},
	})
	//templates["main.tpl"].Execute(os.Stdout, result)
	templates["main.tpl"].Execute(w, result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHome)
	http.ListenAndServe(":6688", mux)
}
