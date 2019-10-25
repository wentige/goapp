package email

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
)

// Setup SMTP Info before Sending email
//
// email.Config.Username = "***"
// email.Config.Password = "***"
// email.Config.Hostname = "***"
// email.Config.Port     = "***"
// email.Config.Form     = "***"
//
// email.SendEmail("to@someone.com", "Title", "Content")

var Config SMTPInfo

type SMTPInfo struct {
	Username string
	Password string
	Hostname string
	Port     int
	From     string
}

// SendEmail sends an email
func SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", Config.Username, Config.Password, Config.Hostname)

	header := make(map[string]string)
	header["From"] = Config.From
	header["To"] = to
	header["Subject"] = subject
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = `text/plain; charset="utf-8"`
	header["Content-Transfer-Encoding"] = "base64"

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	// Send the email
	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", Config.Hostname, Config.Port),
		auth,
		Config.From,
		[]string{to},
		[]byte(message),
	)

	return err
}
