//http://tour.golang.org/#7
package main

import "fmt"

func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(2, 3))
}
