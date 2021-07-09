package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	request2 "wangy1.top/my_cloud/internal/handler/user/request"
	"wangy1.top/my_cloud/internal/model"
)

type Handler struct {
}

var userQuery = model.UserQuery{}

func (uc *Handler) List(c *gin.Context) {
	params := &request2.UserListParams{}
	err := c.BindJSON(params)
	if err != nil {
		return
	}
	log.Debugf("list params:%v", params)
	list, err := userQuery.List(params)
	if err != nil {
		return
	}
	c.JSON(200, list)
}

func (uc *Handler) Get(c *gin.Context) {
	params := c.Params
	fmt.Println(params)

	c.JSON(200, gin.H{
		"code":    200,
		"message": "get success!",
		"data":    params,
	})
}

func (uc *Handler) Post(c *gin.Context) {
	var req request2.UserParams
	params := c.ShouldBindJSON(&req)
	log.Infof("参数：", params)
	//db.DB.Create(user)
	c.JSON(200, gin.H{
		"code":    200,
		"message": "post success!",
		"data":    req,
	})

}

func (uc *Handler) Put(c *gin.Context) {

}

func (uc *Handler) Delete(c *gin.Context) {

}
