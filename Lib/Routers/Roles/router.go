package Roles

import (
	"github.com/gin-gonic/gin"
	"mygo/Controller"
)

func Routers(e *gin.Engine) {
	r := new(Controller.RoleController)
	e.GET("/insert_role", r.WriteData)
}
