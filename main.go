package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type Request struct {
	From   string
	To     string
	Sender string

	Subject string
	Body    string
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		die(fmt.Errorf("missing argument"))
	}

	event := []byte(os.Args[1])
	var req Request
	if err := json.Unmarshal(event, &req); err != nil {
		die(err)
	}

	sess, err := session.NewSession()
	if err != nil {
		die(err)
	}

	s := ses.New(sess)

	_, err = s.SendEmail(&ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(req.To),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(req.Body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(req.Subject),
			},
		},
		Source: aws.String(req.Sender),
		ReplyToAddresses: []*string{
			aws.String(req.From),
		},
	})

	if err != nil {
		die(err)
	}

}
