package main

import "github.com/pborman/uuid"
import "fmt"

func main() {
	fmt.Println("New:", uuid.New())
	fmt.Println("NodeId:", uuid.NodeID())
	x := uuid.NewRandom()
	fmt.Println(x)
	fmt.Println(uuid.Parse(x.String()))
	fmt.Println(uuid.NewMD5(x, []byte{100}))
	fmt.Println(x.Time())
	fmt.Println(x.URN())
}
