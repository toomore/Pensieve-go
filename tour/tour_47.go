package main

import "fmt"
import "runtime"

func main() {
    os := runtime.GOOS
    fmt.Println(os)
    switch os {

    case "darwin":
        fmt.Println("OS X.")

    case "linux":
        fmt.Println("Linux.")

    default:
        fmt.Println(os)
    }

}
