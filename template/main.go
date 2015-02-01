package main

import "text/template"
import "os"

type Data struct {
	Name string
	Age  string
	Info string
}

func main() {
	var t = template.New("main")
	var result = &Data{Name: "Toomore"}
	t = template.Must(t.ParseGlob("*.tpl"))
	t.ExecuteTemplate(os.Stdout, "main", result)
}
