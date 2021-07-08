package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	request2 "wangy1.top/my_cloud/internal/handler/request"
)

func Get(c *gin.Context) {
	params := c.Params
	fmt.Println(params)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "get success!",
		"data":    params,
	})

}
func Post(c *gin.Context) {
	var req request2.UserParams
	params := c.ShouldBindJSON(&req)
	fmt.Println(params)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "post success!",
		"data":    req,
	})

}
