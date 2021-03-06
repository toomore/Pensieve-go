package main

import "bufio"
import "fmt"
import "io"
import "log"
import "os"
import "path/filepath"
import "regexp"
import "strings"

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

func replaceString(s string) (result string) {
    result = strings.ToUpper(s)
    return
}

func makeStringMaps(s string, result *map[string]int) {
    x := *result
    for _, i := range strings.Fields(s) {
        x[i] += 1
    }
}

func copyFiles(inFiles io.Reader, outFiles io.Writer) (err error) {
    reader := bufio.NewReader(inFiles)
    writer := bufio.NewWriter(outFiles)

    // Save data.
    defer func() {
        err = writer.Flush()
    }()

    eof := false
    reg := regexp.MustCompile("[A-Za-z]+")
    count := make(map[string]int)
    for !eof {
        line, err := reader.ReadString('\n')
        if err != io.EOF {
            line = reg.ReplaceAllStringFunc(line, replaceString)
            writer.WriteString(line)
            makeStringMaps(line, &count)
            fmt.Println(line)
        } else {
            eof = true
        }
    }
    fmt.Println(count)
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

    copyFiles(inFiles, outFiles)
}
