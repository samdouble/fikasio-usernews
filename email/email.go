package email

import (
    "log"
    "os"
    "strconv"
    mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>fikas.io Daily News</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func Send() {
    port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
    if err != nil {
		port = 587
	}
    server := mail.NewSMTPClient()
	server.Host = os.Getenv("SMTP_HOST")
	server.Port = port
	server.Username = os.Getenv("SMTP_USERNAME")
	server.Password = os.Getenv("SMTP_PASSWORD")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom(os.Getenv("EMAIL_SENDER"))
	email.AddTo(os.Getenv("EMAIL_RECIPIENT"))
	email.SetSubject("fikas.io Daily News")
	email.SetBody(mail.TextHTML, htmlBody)
	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}
