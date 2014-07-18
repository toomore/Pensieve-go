package main

import "fmt"

type Person struct {
    name    string
    age     int
}

func main() {
    a := Person{"toomore", 30}
    fmt.Printf("%s(%d)\n", a.name, a.age)
}
