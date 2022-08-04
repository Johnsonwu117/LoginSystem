package change

import (
	tool "LoginSystem/app/models"
	"LoginSystem/app/models/changes"
	db "LoginSystem/database"
	"fmt"
	"net/smtp"
)

var (
	Host     = "smtp.gmail.com:587"
	Username = "jo890117jo890117@gmail.com"
	Password = "*****"
)

func Change(user_input changes.User) changes.User {
	var user changes.User
	incode := (tool.RandSeq1(8))
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)

	if user.Verifycode == user_input.Verifycode {
		if user.Password == user_input.Password {

			//db.DB.Save(&user_input).Table("users").Where("Email = ?", user_input.Email).Save(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			db.DB.Table("users").Where("Email = ?", user_input.Email).Updates(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			auth := smtp.PlainAuth(Host, Username, Password, "smtp.gmail.com")
			to := []string{user_input.Email}
			str := fmt.Sprintf(Mail2(), user_input.Email, user.Email, incode, user_input.Phone, user_input.Introduction)
			msg := []byte(str)
			smtp.SendMail(
				Host,
				auth,
				Username,
				to,
				msg,
			)
		}
	}
	return user_input
}
