//http://tour.golang.org/#3
package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {
    fmt.Println("Go...")
    fmt.Println("Now", time.Now())
    fmt.Println(os.Open("tour_3.go"))
    fmt.Println(net.Dial("tcp", "google.com:http"))
}
