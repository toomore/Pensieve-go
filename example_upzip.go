package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
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
		unpackZippedFile(file)
	}
	return nil
}

func unpackZippedFile(zipFile *zip.File) {
	writer, _ := os.Create(zipFile.Name)
	defer writer.Close()
	reader, _ := zipFile.Open()
	defer reader.Close()
	if _, err := io.Copy(writer, reader); err != nil {
		fmt.Println("Zip File Copy wrong.")
	}
}

func main() {
	unpackZip("/Volumes/RamDisk/gounzip/exzip.zip")
}
