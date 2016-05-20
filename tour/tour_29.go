//http://tour.golang.org/#29
package main

import "fmt"

type Vertex struct {
    x, y int
}

var (
    p = Vertex{1, 2}
    q = &Vertex{1, 2}
    r = Vertex{x: 1}
    s = Vertex{}
)

func main() {
    q.x = 11
    fmt.Println(p, q, r, s)
}
