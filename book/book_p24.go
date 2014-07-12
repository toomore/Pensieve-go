package main

import "fmt"
import "github.com/toomore/Pensieve-Go/book/stacker/stack"

func main() {
    var s stack.Stack
    s.Push(123)
    s.Push("Toomore")
    s.Push([]string{"A", "B", "C"})
    fmt.Println(s)
    fmt.Println("Len:", s.Len())
    fmt.Println("Cap:", s.Cap())
    fmt.Println("Top:", s.Top())
    fmt.Println("Pop:", s.Pop())
    fmt.Println(s)
}
