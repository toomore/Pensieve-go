//http://tour.golang.org/#39
package main

import "fmt"

type Vertex struct {
    x, y float64
}

//var m map[string]Vertex

func main() {
    m := make(map[string]Vertex)
    m["toomore"] = Vertex{10, 12}
    m["home"] = Vertex{1, 2}
    fmt.Println(m)
    fmt.Println(m["toomore"])
    fmt.Println(m["home"])
}
