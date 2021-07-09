package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"wangy1.top/my_cloud/configs"
	"wangy1.top/my_cloud/internal/model"
	"wangy1.top/my_cloud/internal/routes"
)

func main() {

	configs.InitConfig()
	r := gin.Default()
	model.SetupDb()
	//路由注册
	routes.RouterRegister(r)

	err := r.Run("0.0.0.0:" + viper.GetString("App.port"))
	if err != nil {
		fmt.Println("http server start error!")
		return
	}
}
