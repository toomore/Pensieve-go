package main

import "fmt"

type composer struct {
    name string
    age  int
}

// struct pointer
func main() {
    toomore := composer{"toomore", 30}
    fmt.Println(toomore)
    fmt.Println(toomore.name, toomore.age)
    toomore2 := new(composer) //pointer
    toomore2.age = 30
    toomore2.name = "toomore2"
    fmt.Println(toomore2)

    toomore3 := &composer{} //pointer
    toomore3.age = 30
    toomore3.name = "toomore3"
    fmt.Println(toomore3)
    toomore3.age = 31

    fmt.Println(toomore, toomore2, toomore3)
}
