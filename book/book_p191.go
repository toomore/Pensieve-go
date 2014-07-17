package main

import "fmt"

func main() {
    //a := 100
    var a interface{} = 100 // need interface
    b, result := a.(int) // type assertions
    fmt.Println(b, result)
}
