package main

import (
	"fmt"
	_ "mygo/Configs"
	"mygo/Lib/Routers"
	"mygo/Lib/Routers/Common"
	"mygo/Lib/Routers/Goods"
	"mygo/Lib/Routers/Orders"
	"mygo/Lib/Routers/Roles"
)

func main() {
	// 加载多个APP的路由配置
	Routers.Include(Goods.Routers, Orders.Routers, Common.Routers, Roles.Routers)

	// 初始化路由
	r := Routers.Init()
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err: %v \n", err)
	}

}
