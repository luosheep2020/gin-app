package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int         `json:"count"`
	Data  interface{} `json:"data"`
}

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	})
}
func Fail(msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: http.StatusBadRequest,
		Msg:  msg,
		Data: nil,
	})
}
