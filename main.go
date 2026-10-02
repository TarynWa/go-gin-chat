// package main

// import (
// 	"bytes"
// 	"log"
// 	"net/http"
// 	"ten/models"
// 	"ten/routers"
// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// )

// var AppJsonConfig = []byte(`
//   {
//   "app": {
//     "port": "8322",
//     "upload_file_path": "e:\\golang\\www\\go-gin-chat\\tmp_images\\",
//     "cookie_key": "4238uihfieh49r3453kjdfg",
//     "serve_type": "GoServe",
//     "sm_token": "xxxxxxxxxxx",
//     "debug_mod": "true"
//   },
//   "mysql": {
//     "dsn": "admin:admin@tcp(127.0.0.1:3306)/go_gin_chat?charset=utf8mb4&parseTime=True&loc=Local"
//   }
// }
// `)

// func init(){
// 	viper.SetConfigType("json")
// 	if err := viper.ReadConfig(bytes.NewBuffer(AppJsonConfig));err!=nil{
// 		if _,ok :=err.(viper.ConfigFileNotFoundError);ok{
// 			log.Println("不支持相应的配置文件")
// 		} else {
// 			log.Println("读取config配置失败")
// 		}
// 		log.Fatal(err)
// 	}
// 	models.InitDB()
// }

// func main() {
// 	gin.SetMode(gin.ReleaseMode)
// 	port := viper.GetString(`app.port`)
// 	router := routers.Initrouter()
// 	http.ListenAndServe(":"+port,router)
// }

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/",func(c *gin.Context){
		c.JSON(http.StatusOK,gin.H{"message":"hello world"})
	})
	r.GET("/users/:id",func(c *gin.Context){
		id := c.Param("id")
		c.String(http.StatusOK,"USER ID :%s",id)
	})
	r.GET("/static/*filepath", func(c *gin.Context) {
    path := c.Param("filepath")  // /css/app.css
    c.String(200, "File: %s", path)
})
	r.Run(":8088")
}