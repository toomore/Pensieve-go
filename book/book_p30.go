package main

import "fmt"
import "log"
import "os"
import "path/filepath"

func getFilenamesfromCmd() (inFiles, outFiles string, err error) {
    if len(os.Args) < 2 {
        return "", "", fmt.Errorf("Error: %s", filepath.Base(os.Args[0]))
    }

    if len(os.Args) > 1 {
        inFiles = os.Args[1]
        if len(os.Args) > 2 {
            outFiles = os.Args[2]
        }
    }

    if outFiles == "" || inFiles == outFiles {
        log.Fatal("Files fail.")
    }
    return
}

func main() {
    fmt.Println("In book page 30")
    inFiles, outFiles, err := getFilenamesfromCmd()
    fmt.Println(inFiles, outFiles, err)
}
