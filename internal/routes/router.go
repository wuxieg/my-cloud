package routes

import "github.com/gin-gonic/gin"

func RouterRegister(r *gin.Engine) {
	UserRegister(r)
}
