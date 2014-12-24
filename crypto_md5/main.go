package main

import (
	"crypto/md5"
	"fmt"
)

func md5String(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func main() {
	fmt.Println(md5String("Toomore"))
}
