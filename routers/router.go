package routers

import (
	"net/http"
	"ten/static"
	"ten/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(viper.GetString(`app.cookie_key`)))
	return sessions.Sessions("go-gin-chat", store)
}
func Initrouter() *gin.Engine{
	router := gin.New()
	if viper.GetString(`app.debug_mod`) == "false" {
		// live 模式 打包用
		router.StaticFS("/static", http.FS(static.EmbedStatic))
	}else{
		// dev 开发用 避免修改静态资源需要重启服务
		router.StaticFS("/static", http.Dir("static"))
	}
	
	sr := router.Group("/",EnableCookieSession())
	{
		sr.GET("/",controller.Index)
	}
	return router
}