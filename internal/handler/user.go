package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"wangy1.top/my_cloud/internal/handler/request"
)

type UserHandler struct {
}

func (uc *UserHandler) Get(c *gin.Context) {
	params := c.Params
	fmt.Println(params)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "get success!",
		"data":    params,
	})
}

func (uc *UserHandler) Post(c *gin.Context) {
	var req request.UserParams
	params := c.ShouldBindJSON(&req)
	log.Infof("参数：", params)
	//db.DB.Create(user)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "post success!",
		"data":    req,
	})

}

func (uc *UserHandler) Put(c *gin.Context) {

}

func (uc *UserHandler) Delete(c *gin.Context) {

}
