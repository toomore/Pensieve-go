package main

import (
	"encoding/json"
	"fmt"
)

// Person json struct
type Person struct {
	Name  string
	Age   int8
	Skill string `json:"skill,omitempty"`
}

func main() {
	// encode
	//data := Person{Name: "Toomore", Age: 30, Skill: "No"}
	data := Person{Name: "Toomore", Age: 30}
	j, _ := json.Marshal(data)
	fmt.Printf("%s\n", j)

	// decode
	stringData := []byte(`{"name": "Toomore", "Age": 30, "Skill": "NO"}`)
	var jsonBlob Person
	json.Unmarshal(stringData, &jsonBlob)
	fmt.Printf("%+v\n", jsonBlob)
}
