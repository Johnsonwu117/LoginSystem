package invite

import (
	"LoginSystem/app/models/invites"
	db "LoginSystem/database"
	"fmt"
	"net/smtp"
)

var (
	host     = "smtp.gmail.com:587"
	username = "jo890117jo890117@gmail.com"
	password = "*****"
)

func InviteMember(user_input invites.User) invites.User {
	var user invites.User
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
