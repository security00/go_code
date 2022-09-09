package Controller

import (
	"github.com/gin-gonic/gin"
	"mygo/Entities/Requests"
	"net/http"
)

type IndexController struct {

}

func (i *IndexController) Index(c *gin.Context)  {
	var json Requests.Login
	if err := c.ShouldBindJSON(&json); err != nil{
		c.JSON(http.StatusOK,gin.H{"error" : err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"user":json.User,"passwork": json.Pssword})
}

func (i *IndexController) ResponseStruct(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Number = 123
	c.JSON(200, msg)
}

func (i *IndexController) ResponseXml(c *gin.Context) {
	c.XML(200, gin.H{"message": "abc"})
}

func (i *IndexController) ResponseYaml(c *gin.Context) {
	c.YAML(200, gin.H{"name": "zhangsan"})
}

