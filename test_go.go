package main

import "fmt"

type container struct {
    x int
}

func mm(x int, msg *chan container) {
    var sum int
    for i:=0; i<=x; i++ {
        sum += i
    }
    *msg <- container{sum}
}

func main () {
    msg := make(chan container)
    go mm(10, &msg)
    result := make(chan container)
    fmt.Scanln(&result)
}
