package main

import "fmt"

const (
    Read = iota
    Write
    Exec
)

func main() {
    fmt.Println(Read, Write, Exec)
}
