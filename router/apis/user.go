package apis

//變更測試
import (
	"LoginSystem/app/models/change"
	"LoginSystem/app/models/cheak"
	"LoginSystem/app/models/invite"
	"LoginSystem/app/models/login"
	"LoginSystem/app/models/mails"

	"net/http"

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

//狀態(尚未驗證):寄信跟驗證碼
func Snadmail(c *gin.Context) {
	user := mails.User{}
	c.BindJSON(&user)
	mails.CreateFirst(user)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "恭喜註冊成功",
	})

}

//狀態(尚未驗證):寄信跟驗證碼<驗證api>
func Cheakvfcode(c *gin.Context) {
	user_input := confirm.User{}
	c.BindJSON(&user_input)
	confirm.CheakVertify(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "驗證成功",
	})

}

//狀態(尚未註冊):填寫密碼

func Logincode(c *gin.Context) {

	user_input := login.User{}
	c.BindJSON(&user_input)
	login.LoginPassWord(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "登入成功",
	})

}

//狀態(修改資料):寄新的驗證碼信 ，先驗證他的驗證碼
func Changefile(c *gin.Context) {

	user_input := change.User{}
	c.BindJSON(&user_input)
	change.Change(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "修改資料成功",
	})

}

//填寫邀請碼
func Invite(c *gin.Context) {
	user_input := invite.User{}
	c.BindJSON(&user_input)
	invite.InviteMember(user_input)
	c.JSON(http.StatusOK, gin.H{
		"訊息": "填寫邀請碼成功",
	})
}
