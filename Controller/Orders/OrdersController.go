package Orders

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/** 本示例延时函数实现接口的使用 */
type Hader interface {
	S(int, int) int
}

// 提供了一种用时定义的方式
type HaderFunc func(int, int) int

func (f HaderFunc) S(a int, b int) int {
	return f(a, b)
}

type OrdersController struct {
}

func (o *OrdersController) Index(c *gin.Context) {
	var a Hader = HaderFunc(func(i int, i2 int) int {
		return i + i2
	})
	var b Hader = HaderFunc(func(i int, i2 int) int {
		return i - i2
	})
	panic("这是一个错误，请捕获我")
	fmt.Println(a.S(1, 1), b.S(1, 1))
	c.Get("orderId")
	fmt.Println("this is order index")
}
