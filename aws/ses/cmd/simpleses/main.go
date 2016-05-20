package main

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/toomore/Pensieve-go/aws/ses/simpleses"
)

var user = &mail.Address{
	Name:    "講太多",
	Address: "...@gmail.com",
}

var sender = &mail.Address{
	Name:    "講太多",
	Address: "...@gmail.com",
}

func main() {
	ses := simpleses.New(os.Getenv("AWSID"), os.Getenv("AWSKEY"))
	msg := simpleses.Message([]*mail.Address{user}, sender, "This is send from SimpleSES.", "<b>Hello Toomore~.</b>")
	result, err := ses.SendEmail(msg)
	fmt.Printf("%+v [Error: %s]\n", result, err)
}
