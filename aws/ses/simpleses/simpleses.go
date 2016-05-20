package simpleses

import (
	"net/mail"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

// New to new a ses.SES
func New(AWSID, AWSKEY string) *ses.SES {
	var config = session.New(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(AWSID, AWSKEY, ""),
	})
	return ses.New(config)
}

// Message to render a ses.SendEmailInput
func Message(ToUsers []*mail.Address, Sender *mail.Address, Subject, Content string) *ses.SendEmailInput {
	var toUsers []*string
	var mailCharset = aws.String("UTF-8")
	var mailContent = aws.String(Content)
	var mailSubject = aws.String(Subject)

	toUsers = make([]*string, len(ToUsers))
	for i, v := range ToUsers {
		toUsers[i] = aws.String(v.String())
	}

	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: toUsers,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: mailCharset,
					Data:    mailContent,
				},
			},
			Subject: &ses.Content{
				Charset: mailCharset,
				Data:    mailSubject,
			},
		},
		Source: aws.String(Sender.String()),
	}
}
