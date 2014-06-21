package main

import "fmt"

func main() {
    elements := map[string]string{
        "name": "toomore",
    }
    fmt.Println(elements)

    mapelements := map[string]map[string]string{
        "Toomore": map[string]string{
            "name": "Toomore",
            "age": "23",
        },
        "Toomore2": map[string]string{
            "name": "Toomore2",
            "age": "29",
        },
    }
    fmt.Println(mapelements)
}
