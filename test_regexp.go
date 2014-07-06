package main

import "fmt"
import "regexp"
import "strings"

func processString(s string) (result string) {
    result = strings.ToUpper(s)
    return
}

func main() {
    reg := regexp.MustCompile("[A-Za-z]{1}")
    result := reg.ReplaceAllStringFunc("I am Toomore", processString)
    fmt.Println(result)
}
