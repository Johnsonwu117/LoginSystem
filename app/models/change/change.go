package change

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

func randSeq1(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}

	return string(b)
}
func Change(user_input User) User {
	var user User
	incode := (randSeq1(8))
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)

	if user.Verifycode == user_input.Verifycode {
		if user.Password == user_input.Password {

			//db.DB.Save(&user_input).Table("users").Where("Email = ?", user_input.Email).Save(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			db.DB.Table("users").Where("Email = ?", user_input.Email).Updates(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")
			to := []string{user_input.Email}
			str := fmt.Sprintf(Mail2(), user_input.Email, user.Email, incode, user_input.Phone, user_input.Introduction)
			msg := []byte(str)
			smtp.SendMail(
				host,
				auth,
				username,
				to,
				msg,
			)
		}
	}
	return user_input
}
