//http://tour.golang.org/#4
package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(5)
    fmt.Println(rand.Intn(10))
}

