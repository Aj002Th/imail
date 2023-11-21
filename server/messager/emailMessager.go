package messager

import "github.com/Aj002Th/imail/common/email"

type EmailMessager struct {
	sender email.EmailSender
}

func NewEmailMessager() *EmailMessager {
	return &EmailMessager{
		sender: email.NewEmailSender(),
	}
}

func (e *EmailMessager) Push(topic, msg string) error {
	return e.sender.SendEmail(topic, msg)
}
