package Mail

import (
	"log"
	"net/smtp"
	"os"
)

func SendEmail(context string, usermail string, theme string) string {
	from := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	to := usermail
	smtpAddress := os.Getenv("SMTP")
	//формирование сообщения
	message := []byte(
		"From: " + from + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: Password Reset\r\n\r\n" +
			context)

	//проверка безопасности яндекса
	auth := smtp.PlainAuth("", from, password, os.Getenv("HOST"))

	//отправка сообщения
	err := smtp.SendMail(smtpAddress, auth, from, []string{to}, message)
	if err != nil {
		log.Fatal("Произошла ошибка отправки сообщения пользователю [ ", usermail, " ] причина: ", err)
		return "ошибка отправки сообщения"
	}
	log.Println("Сообщение с заголовком [ ", theme, " ] пользователю [ ", usermail, " ]")
	return ""
}
