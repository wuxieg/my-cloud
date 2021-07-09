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
	// 配置读取
	configs.InitConfig()

	// 默认路由
	r := gin.Default()
	// model 对象注入
	model.SetupDb()
	//路由注册
	routes.RouterRegister(r)

	// server start
	err := r.Run("0.0.0.0:" + viper.GetString("App.port"))
	if err != nil {
		fmt.Println("http server start error!")
		return
	}
}
