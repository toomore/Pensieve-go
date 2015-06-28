package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var profile = credentials.NewStaticCredentials(os.Getenv("AWSID"), os.Getenv("AWSKEY"), "")
var config = aws.DefaultConfig

func init() {
	config.Region = "us-east-1"
	config.Credentials = profile
}

func main() {
	//ses := ses.New(aws.)
	fmt.Printf("%+v\n", config)
}
