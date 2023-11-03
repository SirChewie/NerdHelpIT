package models

import (
	"fmt"

	"gopkg.in/mail.v2"
)

const (
	DefaultSender = "support@nerdhelpit.com"
)

type Email struct {
	From      string
	To        string
	Subject   string
	PlainText string
	HTML      string
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
	return &es
}

type EmailService struct {
	DefaultSender string

	// unexported fields
	dialer *mail.Dialer
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	es.setFrom(msg, email)
	msg.SetHeader("Subject", email.Subject)
	switch {
	case email.PlainText != "" && email.HTML != "":
		msg.SetBody("text/plain", email.PlainText)
		msg.AddAlternative("text/html", email.HTML)
	case email.PlainText != "":
		msg.SetBody("text/plain", email.PlainText)
	case email.HTML != "":
		msg.SetBody("text/html", email.HTML)
	}
	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

// Function sends the filled out from to the given email
func (es *EmailService) EmailForm(to, from string) error {
	email := Email{
		Subject:   "New form submission",
		To:        to,                                 // TODO: pull from .ENV
		From:      from,                               // TODO: pull from .ENV
		PlainText: "Hello this is some plain text.",   // TODO: use plain text from form
		HTML:      `<h1>Hello this is some HTML</h1>`, // TODO: make the form into a HTML template
	}
	err := es.Send(email)
	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}
	return nil
}

func (es *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}
