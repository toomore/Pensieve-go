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
		fmt.Printf("%v %s\t%d\t%d\t%.2f%%\n",
			mode, file.Name, file.CompressedSize, file.UncompressedSize,
			(float64(file.UncompressedSize)-float64(file.CompressedSize))/float64(file.UncompressedSize)*100)
	}
	return nil
}

func unpackZippedFile(filename string, zipFile *zip.File) {
}

func main() {
	unpackZip("/Volumes/RamDisk/gounzip/exzip.zip")
}
