package Goods

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GoodsController struct {
	
}

func (g *GoodsController) Index(c *gin.Context)  {
	goodId := c.DefaultQuery("goodId","0")
	fmt.Println("goodId:",goodId)
	c.JSON(http.StatusOK,gin.H{"goodId":goodId})
}
