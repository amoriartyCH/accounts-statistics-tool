package aws

import (
	"bytes"
	encsv "encoding/csv"
	"github.com/amoriartyCH/accounts-statistics-tool/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/gomail.v2"
	"io"
)

const (
	subject = "Small Full Accounts Statistics"
	body = "<h1>SFA Stats</h1><p>Attached is the CSV of the stats</p>"
)

/*
Generate and send an email using amazon's Golang sdk.
Email contains the CSV statistics report as an attachment.

To change destination emails, simply add or remove them from the destinations []*string.
 */
func GenerateEmail(csv *models.CSV) error {

	sess, err := session.NewSession(&aws.Config{
		Region:aws.String("eu-west-2")},
	)
	if err != nil {
		return err
	}

	// Create an SES session.
	svc := ses.New(sess)

	source := aws.String("XXX <xxx@xxx.com>")
	destinations := []*string{aws.String("xxx <xxx@xxx.com>")}

	msg := gomail.NewMessage()
	msg.SetHeader("From", "alex@example.com")
	msg.SetHeader("To", "bob@example.com", "cora@example.com")
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	msg.Attach(csv.FileName, gomail.SetCopyFunc(func(w io.Writer) error {

		writer := encsv.NewWriter(w)
		err := writer.WriteAll(csv.Data.ToCSV()) // converts the csv data to a byte array and dumps it to `w`
		return err
	}))

	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	message := ses.RawMessage{ Data: emailRaw.Bytes() }

	input := ses.SendRawEmailInput{Source: source, Destinations: destinations, RawMessage: &message}
	_, err = svc.SendRawEmail(&input)
	if err != nil {
		return err
	}

	return nil
}
