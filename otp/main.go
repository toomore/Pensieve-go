package main

import (
	"bytes"
	"flag"
	"image/png"
	"io/ioutil"
	"log"

	"github.com/pquerna/otp/totp"
)

func gen() {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "toomore.net",
		AccountName: "toomore",
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(key.Secret())

	var buf bytes.Buffer
	img, err := key.Image(200, 200)
	if err != nil {
		log.Fatalln(err)
	}
	png.Encode(&buf, img)
	err = ioutil.WriteFile("qr.png", buf.Bytes(), 0600)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Save OTP QRcode to `qr.png`.")
	}
}

func valid(otpKey string, otp string) bool {
	return totp.Validate(otp, otpKey)
}

func main() {
	flag.Parse()
	if flag.Arg(0) == "gen" {
		gen()
	} else {
		log.Println(valid(flag.Arg(0), flag.Arg(1)))
	}
}
