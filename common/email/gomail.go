package email

import "gopkg.in/gomail.v2"

func test() {
	m := gomail.NewMessage()
	m.SetHeader("From", "2581407059@qq.com")
	m.SetHeader("To", "bob@example.com", "2581407059@qq.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.example.com", 587, "user", "123456")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
