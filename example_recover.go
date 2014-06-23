package main

import "fmt"

func main() {
    defer func() {
        str := recover()
        fmt.Println("panic:", str)
    }()
    panic("Warning")
}
