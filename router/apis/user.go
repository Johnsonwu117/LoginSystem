package apis

//變更測試
import (
	db "LoginSystem/database"
	"LoginSystem/models/mails"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
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

func randSeq(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}

	return string(b)
}
func randSeq1(n int) string {
	b := make([]rune, n)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = letters[r.Intn(62)]
	}

	return string(b)
}

//1.狀態(尚未驗證):寄信跟驗證碼
func Snadmail(c *gin.Context) {
	user := mails.User{}
	c.BindJSON(&user)
	mails.CreateFirst(user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "恭喜註冊成功",
	})

}

//1-2.狀態(尚未驗證):寄信跟驗證碼<驗證api>
func Cheakvfcode(c *gin.Context) {
	var user, user_input User
	c.BindJSON(&user_input)
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)
	if user.Verifycode == user_input.Verifycode {
		db.DB.Save(&user_input).Table("users").Save(map[string]interface{}{"State": "驗證完成"})
		c.JSON(http.StatusOK, gin.H{
			"message": "已成功驗證!",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "已失敗驗證!",
		})
	}

}

//2.狀態(尚未註冊):填寫密碼

func Logincode(c *gin.Context) {

	var user, user_input User
	c.BindJSON(&user_input)
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)

	if user.Verifycode == user_input.Verifycode {
		if user_input.Password == user_input.Twopassword {
			c.JSON(http.StatusOK, gin.H{
				"message": &user_input,
			})
			db.DB.Save(&user_input).Table("users").Save(map[string]interface{}{"State": "登陸完成"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "兩次密碼輸入不一樣!",
			})
		}

	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "驗證碼錯誤哦!",
		})
	}

}

//3.狀態(修改資料):寄新的驗證碼信 ，先驗證他的驗證碼
func Changefile(c *gin.Context) {

	var user, user_input User
	var (
		host     = "smtp.gmail.com:587"
		username = "jo890117jo890117@gmail.com"
		password = "jo890117"
	)
	c.BindJSON(&user_input)
	incode := (randSeq1(8))
	db.DB.Table("users").Where("Email = ?", user_input.Email).Find(&user)

	if user.Verifycode == user_input.Verifycode {
		if user.Password == user_input.Password {
			c.JSON(http.StatusOK, gin.H{
				"message": &user_input,
			})
			//db.DB.Save(&user_input).Table("users").Where("Email = ?", user_input.Email).Save(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			db.DB.Table("users").Where("Email = ?", user_input.Email).Updates(map[string]interface{}{"introduction": user_input.Introduction, "Invitecode": incode, "phone": user_input.Phone})
			auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")
			to := []string{user_input.Email}
			str := fmt.Sprintf("From:jo890117jo890117@gmail.com\r\nTo:%v\r\nSubject:個人資料已經更改\r\n\r\n親愛的會員你好:\r\n\r\n以下為你的個人資料:\r\n\r\n這是你的信箱:%v\r\n\r\n這是你的邀請碼:%08v\r\n\r\n這是你的電話:%v\r\n\r\n這是你的自我介紹:%v\r\n", user_input.Email, user.Email, incode, user_input.Phone, user_input.Introduction)
			msg := []byte(str)
			smtp.SendMail(
				host,
				auth,
				username,
				to,
				msg,
			)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "密碼輸入不一樣!",
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "查無此信箱哦!",
		})
	}

}

//4.填寫邀請碼
func Invite(c *gin.Context) {
	var user, user_input User
	var (
		host     = "smtp.gmail.com:587"
		username = "jo890117jo890117@gmail.com"
		password = "jo890117"
	)
	c.BindJSON(&user_input)
	db.DB.Table("users").Where(" Invitecode= ?", user_input.Invitecode).Find(&user)
	if user.Invitecode == user_input.Invitecode {

		user.Invitenum += 1
		db.DB.Save(&user)
		auth := smtp.PlainAuth(host, username, password, "smtp.gmail.com")

		to := []string{user.Email}

		str := fmt.Sprintf("From:jo890117jo890117@gmail.com\r\nTo:%v\r\nSubject:恭喜你邀請會員!!\r\n\r\n親愛的會員你好:\r\n\r\n\r\n以下為你邀請的人數:%v\r\n", user.Email, user.Invitenum)
		msg := []byte(str)
		smtp.SendMail(
			host,
			auth,
			username,
			to,
			msg,
		)
		c.JSON(http.StatusOK, gin.H{
			"message": "邀請人數+1!!",
		})

	}
}
