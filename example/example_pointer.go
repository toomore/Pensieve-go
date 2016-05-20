package main

import "fmt"

func zero(x *int) {
    *x = 0
    fmt.Println("In zero", x)
}

func square(y *float64) {
    *y = *y * *y
}

func main() {
    x := 54
    zero(&x)
    fmt.Println(x)
    fmt.Println(&x)

    y := 1.5
    square(&y)
    fmt.Println(y)
}
