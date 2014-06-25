package main

import "fmt"

func f(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(i)
    }
}

func main() {
    go f(10)
    var input string
    fmt.Scanln(&input)
}
