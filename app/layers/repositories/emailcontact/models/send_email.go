package models

import (
	"fmt"
	"strings"

	"gitlab.com/chaihanij/evat/app/entities"
	"gopkg.in/gomail.v2"
)

type Email struct {
	Sender   string   `form:"sender"`
	Receiver []string `form:"receiver"`
	CC       []string `form:"cc"`
	Subject  string   `form:"subject"`
	Content  string   `form:"content"`
	Password string
}

func SendEmail(data *entities.CreateContactEmail) bool {

	var (
		email = Email{}
	)
	m := gomail.NewMessage()

	subject := data.Title
	content := data.Description
	receiver := data.Email
	cc := ""

	// payload := &bytes.Buffer{}
	// writer := multipart.NewWriter(payload)
	// _ = writer.WriteField("subject", subject)
	// _ = writer.WriteField("content", content)
	// _ = writer.WriteField("receiver", receiver)
	// _ = writer.WriteField("cc", cc)
	// err := writer.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	email.Sender = "sanch_ai@hotmail.com"
	email.Password = "0877380568"

	mailto := strings.Split(receiver, ",")
	mailcc := strings.Split(cc, ",")

	receiverto := make([]string, len(mailto))
	for i, recipient := range mailto {
		receiverto[i] = m.FormatAddress(recipient, "")
	}

	addresses := make([]string, len(mailcc))
	for i, recipient := range mailcc {
		addresses[i] = m.FormatAddress(recipient, "")
	}

	m.SetHeader("To", receiverto...)
	m.SetHeader("From", email.Sender)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	if cc != "" {
		m.SetHeader("Cc", addresses...)
	}
	d := gomail.NewDialer("smtp.office365.com", 587, email.Sender, email.Password)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("error", err)
		return false
	} else {
		fmt.Println("success -- ")
		return true

	}

}
