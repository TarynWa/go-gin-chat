package controller

import (
	"net/http"
	"ten/ws/primary"
	"ten/services/user_service"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context){
	userInfo := user_service.GetUserInfo(c)
	if len(userInfo) > 0 {
		c.Redirect(http.StatusFound, "/home")
		return
	}
	OnlineUserCount := primary.OnlineUserCount()

	c.HTML(http.StatusOK, "login.html", gin.H{
		"OnlineUserCount": OnlineUserCount,
	})
}