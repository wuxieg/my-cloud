package routes

import (
	"github.com/gin-gonic/gin"
	"wangy1.top/my_cloud/internal/handler/user"
)

func UserRegister(r *gin.Engine) {
	userHandler := user.Handler{}
	users := r.Group("/users")
	{
		users.GET("/", userHandler.List)
		users.GET("/:id", userHandler.Get)
		users.POST("/", userHandler.Post)
		users.PUT("/", userHandler.Put)
		users.DELETE("/", userHandler.Delete)

	}

}
