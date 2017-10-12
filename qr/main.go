package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

var (
	msg          = flag.String("f", "./msg.txt", "message file")
	size         = flag.Int("s", 600, "PNG size")
	q            = flag.String("q", "M", "L(7%), M(15%), Q(25%), H(30%)")
	outputbase64 = flag.Bool("b64", false, "Output base64 file")

	qm = map[string]qr.ErrorCorrectionLevel{
		"L": qr.L,
		"M": qr.M,
		"Q": qr.Q,
		"H": qr.H,
	}
)

func main() {
	flag.Parse()

	content, err := ioutil.ReadFile(*msg)
	if err != nil {
		log.Fatal(err)
	}
	// Create the barcode
	qrCode, _ := qr.Encode(string(content), qm[*q], qr.Unicode)

	// Scale the barcode to size*size pixels
	qrCode, _ = barcode.Scale(qrCode, *size, *size)

	// create the output file
	now := time.Now().Unix()
	file, _ := os.Create(fmt.Sprintf("qrcode_%d.png", now))
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)

	// convert to base64
	if *outputbase64 {
		b64file, _ := os.Create(fmt.Sprintf("qrcode_%d_base64.txt", now))
		defer b64file.Close()

		buf := new(bytes.Buffer)
		png.Encode(buf, qrCode)

		b64file.Write([]byte(fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buf.Bytes()))))
	}
}
