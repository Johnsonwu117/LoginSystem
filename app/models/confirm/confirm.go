package confirm

import (
	db "LoginSystem/database"
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

func CheakVertify(user_input User) User {
	var user User
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)
	if user.Verifycode == user_input.Verifycode {
		db.DB.Save(&user_input).Table("users").Save(map[string]interface{}{"State": "驗證完成"})
	}
	return user_input
}

// marge test
// 看發生什麼事情
