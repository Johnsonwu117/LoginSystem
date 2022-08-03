package routers

import (
	"LoginSystem/apis"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/tset")
	user.POST("/register", apis.Snadmail)
	user.POST("/confirm", apis.Cheakvfcode)
	user.POST("/login", apis.Logincode)
	user.POST("/change", apis.Changefile)
	user.POST("/invite", apis.Invite)
	//user.PUT("/:UserName", apis.UpdateList)
	//user.PUT("/:UserName", apis.UpdateList)
}
