package main

import (
	"html/template"
	"log"
	"os"
)

type data struct {
	Name string
	Age  int64
}

func main() {
	toomore := &data{Name: "Toomore", Age: 30}

	if t, err := template.ParseFiles("./content.txt"); err == nil {
		t.Execute(os.Stdout, toomore)
	} else {
		log.Println(err)
	}
}
