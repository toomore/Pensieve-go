package main

import (
	"fmt"
	"net/mail"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
)

var user = &mail.Address{
	Name:    "講太多",
	Address: "toomore0929@gmail.com",
}

var sender = &mail.Address{
	Name:    "講太多",
	Address: "me@toomore.net",
}

func message() *ses.SendEmailInput {
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(user.String())},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				HTML: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String("Hello Toomore!"),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String("Test send from aws Go SDK"),
			},
		},
		Source: aws.String(sender.String()),
	}
}

func main() {
	var profile = credentials.NewStaticCredentials(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "")
	var config = aws.DefaultConfig
	config.Region = "us-east-1"
	config.Credentials = profile
	msg := ses.New(config)
	fmt.Printf("%+v\n", msg)
	fmt.Println(msg.SendEmail(message()))
}
