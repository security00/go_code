package Orders

import (
	"github.com/gin-gonic/gin"
	. "mygo/Controller/Orders"
)

func Routers(e *gin.Engine) {
	o := new(OrdersController)
	e.GET("/order",o.Index)
}
