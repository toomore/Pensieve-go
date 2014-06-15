//http://tour.golang.org/#28
package main

import "fmt"

type Vertex struct {
    x int
    y int
}

func main() {
    v := Vertex{1, 2}
    q := &v //pointer
    q.x = 1e9
    v.y = 5
    fmt.Println(v, q)
}
