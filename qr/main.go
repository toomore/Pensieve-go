package main

import (
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func main() {
	content, err := ioutil.ReadFile("./msg.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Create the barcode
	qrCode, _ := qr.Encode(string(content), qr.Q, qr.Unicode)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, 200, 200)

	// create the output file
	file, _ := os.Create("qrcode.png")
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}
