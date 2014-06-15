//http://tour.golang.org/#30
package main

import "fmt"

type Vertex struct {
    x, y int
}

func main() {
    v := new(Vertex)
    fmt.Println(v)
    v.x, v.y = 1, 2
    fmt.Println(v)
}
