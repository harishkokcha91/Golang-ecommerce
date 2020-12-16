package models

import (
	fmt "fmt"
	"net/smtp"
)

const (
	fromAddress       = "useremail"
	fromEmailPassword = "password"
	smtpServer        = "smtp.gmail.com"
	smptPort          = "587"
)

// SendEmail will send email to given address
func SendEmail(message string, toAddress []string) (response bool, err error) {

	var auth = smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)
	fmt.Println(message)
	var msg = "To: " + toAddress[0] + "\r\n" + "Subject: Website enquiry!\r\n" + "\r\n" + message
	fmt.Println(msg)
	err = smtp.SendMail(smtpServer+":"+smptPort, auth, fromAddress, toAddress, []byte(msg))

	if err == nil {
		return true, nil
	}
	return false, err

}
