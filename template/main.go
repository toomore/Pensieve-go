package main

import (
	"flag"
	"html/template"
	"io"
	"log"
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
	log.Print(req)
	//templates["main.tpl"].Execute(os.Stdout, result)
	templates["main.tpl"].Execute(w, nil)
}

func serveUser(w http.ResponseWriter, req *http.Request) {
	log.Print(req)
	var result = &Data{Name: req.URL.Path}
	//templates["main.tpl"].Execute(os.Stdout, result)
	templates["user.tpl"].Execute(w, result)
}

var httpAttr = flag.String("http", ":6688", "Http and port.")

func main() {
	flag.Parse()

	parseTemplate([][]string{
		{"main.tpl", "layout.tpl"},
		{"user.tpl", "layout.tpl"},
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/", serveHome)
	mux.HandleFunc("/user/", serveUser)

	log.Println("Starting...", *httpAttr)

	if err := http.ListenAndServe(*httpAttr, mux); err != nil {
		log.Fatal(err)
	}
}
