package gateway

import (
	"encoding/base64"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
	"unibee/internal/logic/email/sender"
	"unibee/utility"
)

func SendEmailToUser(f *sender.Sender, emailGatewayKey string, mailTo string, subject string, body string) (result string, err error) {
	if f == nil {
		f = sender.GetDefaultSender()
	}
	from := mail.NewEmail(f.Name, f.Address)
	to := mail.NewEmail(mailTo, mailTo)
	plainTextContent := body
	htmlContent := "<div>" + body + " </div>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(emailGatewayKey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Printf("SendEmailToUser error:%s\n", err.Error())
		return "", err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return utility.MarshalToJsonString(response), nil
}

func SendPdfAttachEmailToUser(f *sender.Sender, emailGatewayKey string, mailTo string, subject string, body string, pdfFilePath string, pdfFileName string) (result string, err error) {
	if f == nil {
		f = sender.GetDefaultSender()
	}
	from := mail.NewEmail(f.Name, f.Address)
	to := mail.NewEmail(mailTo, mailTo)
	plainTextContent := body
	htmlContent := "<div>" + body + " </div>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	attach := mail.NewAttachment()
	dat, err := os.ReadFile(pdfFilePath)
	if err != nil {
		fmt.Println(err)
	}
	encoded := base64.StdEncoding.EncodeToString(dat)
	attach.SetContent(encoded)
	attach.SetType("application/pdf")
	attach.SetFilename(pdfFileName)
	attach.SetDisposition("attachment")
	message.AddAttachment(attach)
	client := sendgrid.NewSendClient(emailGatewayKey)
	response, err := client.Send(message)
	if err != nil {
		fmt.Printf("SendPdfAttachEmailToUser error:%s\n", err.Error())
		return "", err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	return utility.MarshalToJsonString(response), nil
}
