package email

var EmailSenderImpl EmailSender

type EmailSender interface {
	SendEmail(subject, body string) error
}

func Init() {
	EmailSenderImpl = NewGomail()
}

func NewEmailSender() EmailSender {
	return NewGomail()
}

func SendEmail(subject, body string) error {
	return EmailSenderImpl.SendEmail(subject, body)
}
