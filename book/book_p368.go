package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`[{"name": "Toomore"}]`)
	//data := make(map[string]string)
	//data["name"] = "Toomore"
	type JsonData struct {
		name string
	}
	json_data := []JsonData{}
	json.Unmarshal(data, &json_data)
	fmt.Printf("%+v", json_data)
}
