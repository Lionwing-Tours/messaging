package mail

/*==========================Lionwing Tours (Go) API====================================

	© 2016 Lionwing Tours (Pvt) Ltd. All rights reserved.

	Author(s):
		Surath Jayetileke

==============================================================================*/

import (
	"log"

	"gopkg.in/gomail.v2"
)

const (
	//https://app.mailjet.com
	server   = `in-v3.mailjet.com`
	port     = 25 // OR 587
	username = "00d0ae65b2a04cd5c00322bac4be514f"
	password = "30e7c6925b054cf1d38427bc17c494c0"
)

func SendEmail(recipientAddress, recipientName, senderName, subject, contentHTML, contentText string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", senderName+" <"+"Lionwing Tours"+">")
	msg.SetHeader("To", recipientName+" <"+recipientAddress+">")
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", contentText)
	msg.AddAlternative("text/html", contentHTML)

	dialer := gomail.NewPlainDialer(server, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		log.Println("ERROR sending email:", err)
		return err
	}
	return nil
}
