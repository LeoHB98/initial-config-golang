package utils

import (
	"fmt"
	"strconv"
	"time"

	"gitalpha.com/go/nome_projeto/config"
	"gitalpha.com/go/nome_projeto/models"
	"gopkg.in/gomail.v2"
)

func SendEmail(errors []models.Log, mail *config.Email) {

	header := "Log de Erros"
	message := fmt.Sprintf("%v", errors)
	errorDate := "<br>Data: " + time.Now().Format("2006/01/02 15:04:05")

	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.From)
	msg.SetHeader("To", mail.TO)
	msg.SetHeader("Subject", header)
	msg.SetBody("text/html", message+"<br>"+errorDate)

	port, err := strconv.Atoi(mail.Port)
	if err == nil {
		n := gomail.NewDialer(mail.SMTP, port, mail.From, mail.Pass)

		n.DialAndSend(msg)
	}
}
