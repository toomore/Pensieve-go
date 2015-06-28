package simpleses

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SimpleSES struct {
	awsid  string
	awskey string
}

func (s *SimpleSES) New(AWSID, AWSKEY string) *ses.SES {
	var config = aws.DefaultConfig
	config.Region = "us-east-1"
	config.Credentials = credentials.NewStaticCredentials(AWSID, AWSKEY, "")
	return ses.New(config)
}

func Message(users []*mail.Address, sender *mail.Address, subject, content string) *ses.SendEmailInput {
	var toUsers []*string
	var mailCharset = aws.String("UTF-8")
	var mailContent = aws.String(content)
	var mailSubject = aws.String(subject)

	toUsers = make([]*string, len(users))
	for i, v := range users {
		toUsers[i] = aws.String(v.String())
	}

	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toUsers,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				HTML: &ses.Content{
					Charset: mailCharset,
					Data:    mailContent,
				},
			},
			Subject: &ses.Content{
				Charset: mailCharset,
				Data:    mailSubject,
			},
		},
		Source: aws.String(sender.String()),
	}
}
