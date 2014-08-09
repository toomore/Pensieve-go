package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//data := "{\"name\": \"Toomore\"}"
	data := make(map[string]string)
	data["name"] = "Toomore"
	fmt.Println(json.Marshal(data))
}
