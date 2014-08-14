package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int8
	Skill string
}

func main() {
	data := Person{Name: "Toomore", Age: 30, Skill: "No"}
	j, _ := json.Marshal(data)
	fmt.Printf("%s\n", j)
}
