package Goods

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GoodMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context){
		fmt.Println("执行了good组中间件")
		c.Next()
	}
}