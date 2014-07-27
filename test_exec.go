package main

import "os/exec"
import "bytes"
import "fmt"

func main() {
	cmd := exec.Command("date")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()
	fmt.Println(out.String())
}
