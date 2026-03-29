package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserRequest struct {
	Username string `json:"user" form:"user" xml:"user" binding:"required"`
	Message  string `json:"msg" form:"msg" xml:"msg"`
}

func main() {
	r := SetupRouter()
	r.Run()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/submit", func(c *gin.Context) {
		var data UserRequest
		//ShouldBind 会根据Content-Type自动选择解析器
		//支持JSON，XML，Form
		err := c.ShouldBind(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errot": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"received_user": data.Username,
			"received_msg":  data.Message,
			"content_type":  c.ContentType(),
		})
	})
	return r
}
