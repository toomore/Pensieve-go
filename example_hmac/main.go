package main

import (
	"crypto/hmac"
	"log"
)
import "golang.org/x/crypto/sha3"

const hashkey = "toomore.net"

func check(key, msg []byte) bool {
	return hmac.Equal(key, hmac.New(sha3.New512, []byte(hashkey)).Sum(msg))
}

func main() {
	h := hmac.New(sha3.New512, []byte(hashkey))
	msg := []byte("Toomore is me")
	encodeKey := h.Sum(msg)
	log.Printf("%x", encodeKey)

	// true
	log.Println(check(encodeKey, msg))

	// false
	log.Println(check([]byte(hashkey), msg))
}
