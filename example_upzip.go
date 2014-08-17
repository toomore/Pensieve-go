package main

import (
	"archive/zip"
	"fmt"
)

func unpackZip(filename string) error {
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		mode := file.Mode()
		fmt.Println(mode, file.FileInfo())
	}
	return nil
}

func main() {
	unpackZip("/Volumes/RamDisk/gounzip/exzip.zip")
}
