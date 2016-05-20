//http://tour.golang.org/#26
package main

import "fmt"

type Vertex struct {
    x int
    y int
}

func main() {
    fmt.Println(Vertex{1, 2})
}
