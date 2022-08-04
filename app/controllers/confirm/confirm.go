package confirm

import (
	db "LoginSystem/database"
	"LoginSystem/app/models/confirms"
)



func ConfirmVertify(user_input confirms.User) confirms.User {
	var user confirms.User
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)
	if user.Verifycode == user_input.Verifycode {
		db.DB.Save(&user_input).Table("users").Save(map[string]interface{}{"State": "驗證完成"})
	}
	return user_input
}

// marge test
// 看發生什麼事情
