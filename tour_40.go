package main

import "fmt"

type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Taiwan": Vertex{1, 2},
    "Japan": Vertex{3, 4},
}

func main() {
    fmt.Println(m)
    fmt.Println(len(m))
    fmt.Println(m["Taiwan"])
}
