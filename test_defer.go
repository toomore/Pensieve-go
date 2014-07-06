package main

import "fmt"

func testDefer(x string) (result string) {
    result = "testDefer result"
    defer fmt.Println("In testDefer defer", x)
    fmt.Println("Before return")
    x += "<edited>" //not works.
    return
}


func main() {
    fmt.Println("Start testDefer")
    fmt.Println(testDefer("Toomore"))
}
