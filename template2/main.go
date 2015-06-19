package main

import (
	"html/template"
	"os"
)

type data struct {
	Name string
	Age  int64
}

func main() {
	toomore := &data{Name: "Toomore", Age: 30}
	t, _ := template.New("tt").Parse("Hi {{.Name}}, your age is {{.Age}}.\n")
	t.Execute(os.Stdout, toomore)
}
