package main

import "fmt"

type Person struct {
    name string
}

//method
func (p *Person) talk() {
    fmt.Println("My name is", p.name)
}

type Android struct {
    Person
    Model string
}

func main() {
    toomore := new(Android)
    toomore.name = "Toomore"
    toomore.talk()
}
