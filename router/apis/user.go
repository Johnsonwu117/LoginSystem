package apis

//變更測試
import (
	"LoginSystem/app/controllers/change"
	"LoginSystem/app/controllers/confirm"
	"LoginSystem/app/controllers/invite"
	"LoginSystem/app/controllers/login"
	"LoginSystem/app/controllers/mails"
	"LoginSystem/app/models/changes"
	"LoginSystem/app/models/confirms"
	"LoginSystem/app/models/invites"
	"LoginSystem/app/models/logins"
	"LoginSystem/app/models/mail"

	"net/http"

	"github.com/gin-gonic/gin"
)

//狀態(尚未驗證):寄信跟驗證碼
func Snadmail(c *gin.Context) {
	user := mail.User{}
	c.BindJSON(&user)
	mails.CreateFirst(user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "恭喜註冊成功",
	})

}

//狀態(尚未驗證):寄信跟驗證碼<驗證api>
func Cheakvfcode(c *gin.Context) {
	user_input := confirms.User{}
	c.BindJSON(&user_input)
	confirm.ConfirmVertify(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "驗證成功",
	})

}

//狀態(尚未註冊):填寫密碼

func Logincode(c *gin.Context) {

	user_input := logins.User{}
	c.BindJSON(&user_input)
	login.LoginPassWord(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "登入成功",
	})

}

//狀態(修改資料):寄新的驗證碼信 ，先驗證他的驗證碼
func Changefile(c *gin.Context) {

	user_input := changes.User{}
	c.BindJSON(&user_input)
	change.Change(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "修改資料成功",
	})

}

//填寫邀請碼
func Invite(c *gin.Context) {
	user_input := invites.User{}
	c.BindJSON(&user_input)
	invite.InviteMember(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "填寫邀請碼成功",
	})
}
