package login

import (
	"LoginSystem/app/models/logins"
	db "LoginSystem/database"
)

func LoginPassWord(user_input logins.User) logins.User {
	var user logins.User
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)

	if user.Verifycode == user_input.Verifycode {
		if user_input.Password == user_input.Twopassword {

			db.DB.Save(&user_input).Table("users").Save(map[string]interface{}{"State": "登入完成"})
		}
	}
	return user_input
}
