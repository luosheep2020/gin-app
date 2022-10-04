package api1

import (
	"gin-app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FindUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := models.FindUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Msg:  "查询成功",
			Data: user,
		})
	}
}
func CreateUser(c *gin.Context) {

}
func FindUsers(c *gin.Context) {
	users, err := models.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code:  http.StatusOK,
			Msg:   "查询成功",
			Count: len(users),
			Data:  users,
		})
	}
}
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.DeleteUser(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"delete": id})
	}
}
