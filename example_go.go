package main

import "fmt"
import "time"
import "math/rand"

func f(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(i)
        amt := time.Duration(rand.Intn(550))
        time.Sleep(time.Millisecond * amt)
    }
}

func main() {
    for i :=0; i < 10; i++ {
        go f(10)
    }
    var input string
    fmt.Scanln(&input)
}
