package main

import "fmt"
import "os"

func getFilenamesfromCmd() (inFiles, outFiles string, err error) {
    return os.Args[1], os.Args[2], nil
}

func main() {
    fmt.Println("In book page 30")
    inFiles, outFiles, err := getFilenamesfromCmd()
    fmt.Println(inFiles, outFiles, err)
}
