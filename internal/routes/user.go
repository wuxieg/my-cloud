package routes

import (
	"github.com/gin-gonic/gin"
	"wangy1.top/my_cloud/internal/handler"
)

func UserRegister(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", handler.Get)
		users.POST("/", handler.Post)
	}

}
