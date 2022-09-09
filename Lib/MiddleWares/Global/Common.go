package Global

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CommonMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("执行了全局中间件")
		c.Next()
	}
}
