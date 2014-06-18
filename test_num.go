package main

import "fmt"

func main() {
    fmt.Printf("%T\n", int8(-127))
    fmt.Printf("%T\n", uint8(255))
    fmt.Printf("%T\n", 1<<8)
}
