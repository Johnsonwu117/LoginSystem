package mails

import (
	db "LoginSystem/database"
	"fmt"
	"math/rand"
	"net/smtp"
	"time"
)

type User struct {
	Id           int    `json:"UserId" form:"id"`
	Name         string `json:"UserName" form:"name"`
	Password     string `json:"UserPassword" form:"password"`
	Twopassword  string `json:"UserPasswordCheak" form:"twopassword"`
	Email        string `json:"UserEmail" form:"email"`
	Verifycode   string `json:"verifycode"`
	Invitecode   string `json:"invitecode"`
	Invitenum    int    `json:"invitenum"`
	State        string `json:"state"`
	Phone        string `json:"phone"`
	Introduction string `json:"introduction"`
}

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

func CreateFirst(user User) User {
	incode := (randSeq(8))
	db.DB.Create(&user).Table("users").Save(map[string]interface{}{"Invitecode": incode})
	auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")
	to := []string{user.Email}
	str := fmt.Sprintf("From:jo890117jo890117@gmail.com\r\nTo:%v\r\nSubject:信箱驗證碼\r\n\r\n親愛的會員你好:\r\n\r\n感謝你註冊會員，以下為你的邀請碼:\r\n\r\n這是你的邀請碼:%08v\r\n\r\n請多多邀請別人。\r\n", user.Email, incode)
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
