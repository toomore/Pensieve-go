package main

import (
	"bytes"
	"flag"
	"image/png"
	"io/ioutil"
	"log"
	"strings"
	"time"

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

func show(otpKey string) string {
	code, err := totp.GenerateCode(otpKey, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return code
}

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	case "gen":
		gen()
	case "show":
		log.Println(show(strings.ToUpper(flag.Arg(1))))
	default:
		log.Println(valid(flag.Arg(0), flag.Arg(1)))
	}
}
