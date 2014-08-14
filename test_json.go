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
	// encode
	data := Person{Name: "Toomore", Age: 30, Skill: "No"}
	j, _ := json.Marshal(data)
	fmt.Printf("%s\n", j)

	// decode
	string_data := []byte(`{"name": "Toomore", "Age": 30, "Skill": "NO"}`)
	var json_blob Person
	json.Unmarshal(string_data, &json_blob)
	fmt.Printf("%+v\n", json_blob)
}
