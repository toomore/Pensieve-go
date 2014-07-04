package main

import "fmt"
import "os"
import "path/filepath"

func getFilenamesfromCmd() (inFiles, outFiles string, err error) {
    return filepath.Base(os.Args[1]), filepath.Base(os.Args[2]), nil
}

func main() {
    fmt.Println("In book page 30")
    inFiles, outFiles, err := getFilenamesfromCmd()
    fmt.Println(inFiles, outFiles, err)
}
