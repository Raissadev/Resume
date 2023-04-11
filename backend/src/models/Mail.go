package models

import (
	"net/smtp"
	"os"
)

type Mail struct {
	Host         string
	Port         string
	Subject      string `json:"subject" validate:"min=3,max=40"`
	Message      string `json:"message" validate:"min=3,max=40"`
	Sender       string `json:"email" validate:"min=3,max=40,regexp=^[_A-Za-z0-9+-]+(\\.[_A-Za-z0-9-]+)*@[A-Za-z0-9-]+(\\.[A-Za-z0-9]+)*(\\.[A-Za-z]{2\\,})$"`
	MailReceiver *MailReceiver
}

type MailReceiver struct {
	receiver string
	password string
	mail     string
}

func (m *Mail) New() *Mail {
	return &Mail{
		Host: os.Getenv("SMTP_HOST"),
		Port: os.Getenv("SMTP_PORT"),
		MailReceiver: &MailReceiver{
			mail:     os.Getenv("MAIL_RECEIVER"),
			receiver: os.Getenv("SMTP_USERNAME"),
			password: os.Getenv("SMTP_PASS"),
		},
	}
}

func (m *Mail) Send(params Mail) (Mail, error) {
	s := m.New()

	mail := params

	auth := smtp.PlainAuth("", s.MailReceiver.receiver, s.MailReceiver.password, s.Host)

	msg := `
		To: ` + mail.Sender + `
		From: ` + s.MailReceiver.mail + `
		Subject: ` + mail.Subject + `

		` + mail.Message + `
	`

	err := smtp.SendMail(
		s.Host+":"+s.Port,
		auth, s.MailReceiver.mail, []string{mail.Sender},
		[]byte(msg))

	if err != nil {
		return mail, err
	}

	return mail, nil
}
