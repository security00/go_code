package Goods

import (
	"github.com/gin-gonic/gin"
	"mygo/Controller"
	. "mygo/Controller/Goods"
	"mygo/Lib/MiddleWares/Goods"
)

func Routers(e *gin.Engine) {
	g := new(GoodsController)
	e.GET("/goods",Goods.GoodMiddleWare(),g.Index)
	goc := new(Controller.GoodsForOmoCourseControllor)
	e.POST("/relation",goc.Insert)
}
