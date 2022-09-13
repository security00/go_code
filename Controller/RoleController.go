package Controller

import (
	"github.com/gin-gonic/gin"
	"mygo/Entities/Databases"
	"net/http"
)

type RoleController struct {
}

func (r *RoleController) WriteData(c *gin.Context) {
	ri := new(Databases.Roles)
	_, _ = ri.InsertData(struct{}{})
	c.JSON(http.StatusOK, gin.H{"Msg": "successful"})
}
