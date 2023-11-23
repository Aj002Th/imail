package email

import (
	"github.com/Aj002Th/imail/common/config"
	"gopkg.in/gomail.v2"
)

type Gomail struct {
	dialer *gomail.Dialer
}

func NewGomail() *Gomail {
	sender := config.GetEmailSenderForMessager()
	return &Gomail{
		dialer: gomail.NewDialer(sender.Host, sender.Port, sender.Username, sender.Password),
	}
}

func (g *Gomail) SendEmail(subject, body string) error {
	sender := config.GetEmailSenderForMessager()
	receiver := config.GetEmailReceiversForMessager()

	m := gomail.NewMessage()
	m.SetHeader("From", sender.Nickname+"<"+sender.Username+">")
	m.SetHeader("To", receiver...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return g.dialer.DialAndSend(m)
}
