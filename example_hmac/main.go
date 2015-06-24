package main

import (
	"crypto/hmac"
	"log"
)
import "golang.org/x/crypto/sha3"

const hashkey = "toomore.net"

func main() {
	h := hmac.New(sha3.New512, []byte(hashkey))
	h.Write([]byte("Toomore is me"))
	log.Printf("%x", h.Sum(nil))
}
