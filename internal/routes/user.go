package routes

import (
	"github.com/gin-gonic/gin"
	"wangy1.top/my_cloud/internal/handler"
)

func UserRegister(r *gin.Engine) {
	userHandler := handler.UserHandler{}
	users := r.Group("/users")
	{
		users.GET("/", userHandler.Get)
		users.POST("/", userHandler.Post)
		users.PUT("/", userHandler.Put)
		users.DELETE("/", userHandler.Delete)

	}

}
