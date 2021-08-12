package mail

import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

type GoMailer struct {
	Message *gomail.Message
}

// Method to set Message on GoMailer
func NewMailer() *GoMailer {
	return &GoMailer{
		Message: gomail.NewMessage(),
	}
}

// Method to send email
func (mailer *GoMailer) Send(urlToken string) error {
	mailer.Message.SetHeader("From", "test@example.com")
	mailer.Message.SetHeader("To", "test@example.com")
	mailer.Message.SetHeader("Subject", "Thank you for register!")
	mailer.Message.SetBody("text/html", fmt.Sprintf("<h4>You can continue to register by clicking <a href=\"http://localhost:3000?param=%s\">here</a></h4>", urlToken))

	d := gomail.Dialer{Host: "mailhog", Port: 1025}
	if err := d.DialAndSend(mailer.Message); err != nil {
		return err
	}

	return nil
}
