//http://tour.golang.org/#27
package main

import "fmt"

type Vertex struct {
    x int
    y int
    z bool
}

func main() {
    v := Vertex{1, 2, false}
    v.x = 4
    v.z = true
    fmt.Println(v)
}
