package invite

import (
	db "LoginSystem/database"
	"fmt"
	"net/smtp"
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

var (
	host     = "smtp.gmail.com:587"
	username = "jo890117jo890117@gmail.com"
	password = "jo890117"
)

func InviteMember(user_input User) User {
	var user User
	db.DB.Table("users").Where(" Invitecode= ?", user_input.Invitecode).Find(&user)
	if user.Invitecode == user_input.Invitecode {

		user.Invitenum += 1
		db.DB.Save(&user)
		auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")

		to := []string{user.Email}

		str := fmt.Sprintf(Mail3(), user.Email, user.Invitenum)
		msg := []byte(str)
		smtp.SendMail(
			host,
			auth,
			username,
			to,
			msg,
		)

	}
	return user_input
}
