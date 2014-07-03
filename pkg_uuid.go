package main

import "code.google.com/p/go-uuid/uuid"
import "fmt"

func main() {
    fmt.Println("New:", uuid.New())
    fmt.Println("NodeId:", uuid.NodeID())
    x := uuid.NewRandom()
    fmt.Println(x)
    fmt.Println(uuid.Parse(x.String()))
}
