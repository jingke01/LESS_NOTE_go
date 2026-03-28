# Gin
## Gin路由
### 基本路由
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	r.PUT("/post")
	r.Run(":8080")
}
```
gin中路由是使用httprouter写的 有想法可以自己写一个
### restful风格的API
rest(Representational State Transfer) 表现层状态转化 即 URL定位资源，HTTP描述操作

