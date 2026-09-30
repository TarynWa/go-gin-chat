package user_service

import (
	"ten/services/session"

	"github.com/gin-gonic/gin"
)


func GetUserInfo(c *gin.Context) map[string]interface{} {
	return session.GetSessionUserInfo(c)
}