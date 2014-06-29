package main

import "fmt"

func main() {
    dict := make(map[string]int)
    dict["age"] = 123
    fmt.Println(dict)

    delete(dict, "age")
    fmt.Println(dict)

    fmt.Println(dict["age"])
    ele, err := dict["age"]
    fmt.Println(ele, err)
}
