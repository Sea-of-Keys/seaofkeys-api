package pkg

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(email string, name string, id uint, token string) error {
	emailServer := "smtp.gmail.com"
	emailPort := "587"
	senderEmail := "x.seaofkeys@gmail.com"
	// senderPassword := os.Getenv("EMAILPASS")

	// Recipient email address
	// recipientEmail := "mkronborg7@gmail.com"

	// Compose the email message
	subject := "Set Your Code/Password SeaOfKeys.com"
	// body := `
	// <html>
	// <body>
	// <h1>Hello, ` + name + `</h1>
	// <p>We Need you to set a code to acesse are buildings</p>
	// <b><a href="https://api.seaofkeys.com/web/set/` + token + `">set your code</a></b>
	// </body>
	// </html>
	// `
	body := `
        <html>
            <body>
                <h1>Hello, ` + fmt.Sprintf("%v#%v", name, id) + `</h1>
                <p>We Need you to set a code to acesse are buildings</p>
    <b><a href="https://api.seaofkeys.com/web/token/` + token + `">set your code</a></b>
            </body>
        </html>
    `

	// Convert the email body to a byte slice
	bodyBytes := []byte(body)

	// Create an email message
	emailMessage := "Subject: " + subject + "\r\n" +
		"To: " + email + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	emailMessage = emailMessage + string(bodyBytes)

	// Establish a connection to the SMTP server
	auth := smtp.PlainAuth("", senderEmail, os.Getenv("EMAILPASS"), emailServer)
	err := smtp.SendMail(
		emailServer+":"+emailPort,
		auth,
		senderEmail,
		[]string{email},
		[]byte(emailMessage),
	)
	if err != nil {
		return err
	}

	// Email sent successfully
	println("Email sent successfully")
	return nil
}
