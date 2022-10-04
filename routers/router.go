package routers

import (
	"gin-app/controller/api1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/api/:id", api1.FindUser)
	router.GET("/api/c", api1.CreateUser)
	router.GET("/api/user", api1.FindUsers)
	router.DELETE("/api/:id", api1.DeleteUser)
	return router
}
