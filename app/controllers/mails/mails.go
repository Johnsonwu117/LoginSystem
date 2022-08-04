package mails

import (
	tool "LoginSystem/app/models"
	"LoginSystem/app/models/mail"
	db "LoginSystem/database"
	"fmt"
	"net/smtp"
)

var (
	host     = "smtp.gmail.com:587"
	username = "jo890117jo890117@gmail.com"
	password = "jo890117"
)

func CreateFirst(user mail.User) mail.User {
	incode := (tool.RandSeq(8))
	db.DB.Create(&user).Table("users").Save(map[string]interface{}{"Invitecode": incode})
	auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")
	to := []string{user.Email}
	str := fmt.Sprintf(Mail1(), user.Email, incode)
	msg := []byte(str)
	smtp.SendMail(
		host,
		auth,
		username,
		to,
		msg,
	)
	return user
}
