package Common

import (
	"github.com/gin-gonic/gin"
	. "mygo/Controller"
)

func Routers(e *gin.Engine) {
	c := new(IndexController)
	e.GET("/index",c.Index)
	e.GET("/reponse_struct",c.ResponseStruct)
	e.GET("/reponse_xml",c.ResponseXml)
	e.GET("/reponse_yaml",c.ResponseYaml)
}
