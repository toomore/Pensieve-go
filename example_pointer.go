package main

import "fmt"

func zero(x *int) {
    *x = 0
    fmt.Println("In zero", x)
}

func main() {
    x := 54
    zero(&x)
    fmt.Println(x)
    fmt.Println(&x)
}
