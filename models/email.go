package models

import (
	"fmt"

	"gopkg.in/mail.v2"
)

type Email struct {
	From      string
	To        string
	Subject   string
	PlainText string
	HTML      string
}

type SMTPConfig struct {
	Port     int
	Host     string
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

	// unexported fields
	dialer *mail.Dialer
}

func (es *EmailService) Send(email Email) error {
	msg := mail.NewMessage()
	msg.SetHeader("To", email.To)
	msg.SetHeader("From", email.From)
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

// Function sends the filled out form to the given email
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
