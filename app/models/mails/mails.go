package mails

import (
	"LoginSystem/app/controllers/mail"
	db "LoginSystem/database"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	host     = "smtp.gmail.com:587"
	username = "jo890117jo890117@gmail.com"
	password = "jo890117"
)

func randSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}

	return string(b)
}

func CreateFirst(user mail.User) mail.User {
	incode := (randSeq(8))
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
