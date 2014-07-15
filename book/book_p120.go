package main

import "regexp"
import "fmt"

func main() {
    reg, _ := regexp.Compile(`^[a-z0-9\s]+$`)
    fmt.Println(reg.FindAllString("book 123", -1))
}
