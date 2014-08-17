package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func unpackZip(filename string) error {
	unzipDir := filepath.Dir(filename)
	reader, err := zip.OpenReader(filename)
	if err != nil {
		return err
	}
	defer reader.Close()
	fmt.Printf("%+v", reader)
	for _, file := range reader.File {
		mode := file.Mode()
		fmt.Printf("%v %s\t%d\t%d\t%.2f%%\n",
			mode, file.Name, file.CompressedSize, file.UncompressedSize,
			(float64(file.UncompressedSize)-float64(file.CompressedSize))/float64(file.UncompressedSize)*100)
		unpackZippedFile(file, unzipDir)
	}
	return nil
}

func unpackZippedFile(zipFile *zip.File, unzipDir string) {
	writer, _ := os.Create(filepath.Join(unzipDir, zipFile.Name))
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
