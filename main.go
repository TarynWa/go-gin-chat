package main

import (
	"bytes"
	"log"
	"net/http"
	"ten/models"
	"ten/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)


var AppJsonConfig = []byte(`
  {
  "app": {
    "port": "8322",
    "upload_file_path": "e:\\golang\\www\\go-gin-chat\\tmp_images\\",
    "cookie_key": "4238uihfieh49r3453kjdfg",
    "serve_type": "GoServe",
    "sm_token": "xxxxxxxxxxx",
    "debug_mod": "true"
  },
  "mysql": {
    "dsn": "admin:admin@tcp(127.0.0.1:3306)/go_gin_chat?charset=utf8mb4&parseTime=True&loc=Local"
  }
}  
`)

func init(){
	viper.SetConfigType("json")
	if err := viper.ReadConfig(bytes.NewBuffer(AppJsonConfig));err!=nil{
		if _,ok :=err.(viper.ConfigFileNotFoundError);ok{
			log.Println("不支持相应的配置文件")
		} else {
			log.Println("读取config配置失败")
		}
		log.Fatal(err)
	}
	models.InitDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	port := viper.GetString(`app.port`)
	router := routers.Initrouter()
	http.ListenAndServe(":"+port,router)
}