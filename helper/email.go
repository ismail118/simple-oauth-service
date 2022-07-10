package helper

import "gopkg.in/gomail.v2"

//dialer := gomail.NewDialer(
//CONFIG_SMTP_HOST,
//CONFIG_SMTP_PORT,
//CONFIG_AUTH_EMAIL,
//CONFIG_AUTH_PASSWORD,
//)

type Message struct {
	From        string
	To          []string
	Subject     string
	CC          string
	BodyMessage string
	FilesAttach []string
}

func SendEmail(dialer *gomail.Dialer, message Message) {
	mailer := gomail.NewMessage()
	mailer.SetHeader("Subject", message.Subject)
	mailer.SetHeader("From", dialer.Username)
	mailer.SetHeader("To", message.To...)
	//mailer.SetAddressHeader("Cc", message.CC, "")
	mailer.SetBody("text/plain", message.BodyMessage)
	if message.FilesAttach != nil {
		for _, each := range message.FilesAttach {
			mailer.Attach(each)
		}
	}

	err := dialer.DialAndSend(mailer)
	PanicIfError(err)
}
