package primary

import (
	"ten/ws"
	"ten/ws/go_ws"

	"github.com/spf13/viper"
)

// 定义 serve 的映射关系
var serveMap = map[string]ws.ServeInterface{
	"Serve":   &ws.Serve{},
	"GoServe": &go_ws.GoServe{},
}

func Create() ws.ServeInterface {
	// GoServe or Serve
	_type := viper.GetString("app.serve_type")
	return serveMap[_type]
}
func OnlineUserCount() int {
	return Create().GetOnlineUserCount()
}