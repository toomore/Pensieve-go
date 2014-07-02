package main

import "fmt"
import "stacker/stack"

func main() {
    var s stack.Stack
    s.Push(123)
    s.Push("Toomore")
    fmt.Println(s)
    fmt.Println("Len:", s.Len())
}
