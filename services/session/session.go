package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"ten/models"
)

func GetSessionUserInfo(c *gin.Context)map[string]interface{}{
	session := sessions.Default(c)

	uid := session.Get("uid")

	data := make(map[string]interface{})
	if uid != nil {
		user := models.FindUserByField("id", uid.(string))
		data["uid"] = user.ID
		data["username"] = user.Username
		data["avatar_id"] = user.AvatarId
	}
	return data
}