package main

import (
	"os"
	"text/template"
)

type Data struct {
	Name string
	Age  int
	Info string
}

func main() {
	var t = template.New("")
	var result = &Data{Name: "Toomore"}
	//t = template.Must(t.ParseGlob("*.tpl"))
	t = template.Must(t.ParseFiles([]string{"main.tpl", "layout.tpl"}...))
	t = t.Lookup("Root")
	t.Execute(os.Stdout, result)
}
