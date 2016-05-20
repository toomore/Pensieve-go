package main

import "fmt"

func recursion(x int) int {
    if x == 0 {
        return 1
    }

    return x * recursion(x - 1)
}

func main() {
    fmt.Println(recursion(5))
}
