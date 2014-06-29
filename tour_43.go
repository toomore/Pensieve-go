package main

import "fmt"
import "strings"
import "code.google.com/p/go-tour/wc"

func wordcount(s string) map[string]int {
    result := make(map[string]int)
    for _, v := range strings.Fields(s) {
        result[v] += 1
    }
    return result
}

func main() {
    fmt.Println(strings.Fields("i am toomore, toomore"))
    wc.Test(wordcount)
}
