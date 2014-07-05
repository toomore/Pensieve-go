package main

import "fmt"
import "log"
import "os"
import "path/filepath"

func getFilenamesfromCmd() (inFilename, outFilename string, err error) {
    if len(os.Args) < 2 {
        return "", "", fmt.Errorf("Error: %s", filepath.Base(os.Args[0]))
    }

    if len(os.Args) > 1 {
        inFilename = os.Args[1]
        if len(os.Args) > 2 {
            outFilename = os.Args[2]
        }
    }

    if outFilename == "" || inFilename == outFilename {
        log.Fatal("Files fail.")
    }
    return
}

func main() {
    fmt.Println("In book page 30")
    inFilename, outFilename, err := getFilenamesfromCmd()
    if err != nil {
        fmt.Println(err)
    }

    inFiles, outFiles := os.Stdin, os.Stdout
    if inFiles, err = os.Open(inFilename); err != nil {
        log.Fatal(err)
    }
    defer inFiles.Close()

    if outFiles, err = os.Create(outFilename); err != nil {
        log.Fatal(err)
    }
    defer outFiles.Close()
}
