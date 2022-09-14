package Controller

import (
	"github.com/gin-gonic/gin"
	"mygo/Entities/Requests"
	"mygo/models"
	"net/http"
	"strconv"
)

type IndexController struct {
}

func (i *IndexController) Index(c *gin.Context) {
	var json Requests.Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": json.User, "passwork": json.Pssword})
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

func (i *IndexController) UserName(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Code    int
	}
	t := c.Query("id")
	id, _ := strconv.ParseInt(t, 10, 64)
	u := new(models.User)
	name, err := u.GetName(id)
	if err != nil {
		msg.Name = ""
		msg.Message = "获取数据失败"
		msg.Code = -1000
		c.JSON(200, msg)
		return
	}

	msg.Name = name
	msg.Message = "获取数据成功"
	msg.Code = 1000
	c.JSON(200, msg)
}
