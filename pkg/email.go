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

	subject := "Sæt din kode - SeaOfKeys.com"

	body := `
        <html>
            <body>
                <h1>Hej, ` + fmt.Sprintf("%v#%v", id, name) + `</h1>
                <p>Sæt din kode ved at følge linket</p>
    <b><a href="https://api.seaofkeys.com/web/token/` + token + `">set your code</a></b>
            </body>
        </html>
    `
	bodyBytes := []byte(body)

	emailMessage := "Subject: " + subject + "\r\n" +
		"To: " + email + "\r\n" +
		"MIME-version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n"
	emailMessage = emailMessage + string(bodyBytes)

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

	println("Email sent successfully")
	return nil
}
