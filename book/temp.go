package main

import "fmt"

func main() {
    a := make(map[string]string)
    b := make(map[string]string)
    a["name"] = "Toomore"
    b["age"] = "30"
    c := append(a, b)
    fmt.Println(c)
}
