//http://tour.golang.org/#9
package main

import "fmt"

func swap(x, y string) (string, string) {
    return y, x
}

func main() {
    a, b := swap("world", "Hello")
    fmt.Println(a, b)
}
