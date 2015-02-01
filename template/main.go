package main

import "text/template"
import "os"

func main() {
	var t = template.New("main")
	t = template.Must(t.ParseFiles("main.tpl"))
	t.ExecuteTemplate(os.Stdout, "main", nil)
}
