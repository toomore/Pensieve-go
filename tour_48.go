package main

import "fmt"
import "time"

func main() {
    today := time.Now().Weekday()
    fmt.Printf("%s\n", today+7)
    fmt.Println(time.Now())
    switch time.Saturday {
    case today + 0:
        fmt.Println("Today")
    case today + 1:
        fmt.Println("Tomorrow.")
    case today + 2:
        fmt.Println("In two days")
    default:
        fmt.Println("Too far away.")
    }
}
