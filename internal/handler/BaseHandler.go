package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessResult(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func FailResult(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, err)
}
