package Routers

import (
	"github.com/gin-gonic/gin"
	"mygo/Lib/MiddleWares/Global"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	r.Use(Global.CommonMiddleWare(),Global.Myloger())
	for _, opt := range options {
		opt(r)
	}
	r.LoadHTMLGlob("Views/*tmpl")
	return r
}