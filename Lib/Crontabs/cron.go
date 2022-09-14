package Crontabs

import (
	"fmt"
	"time"
)
import cron "github.com/robfig/cron/v3"

var cr *cron.Cron

func Init() {
	cr := cron.New(cron.WithSeconds())
	// 传统表达式写法: 每秒执行一次
	_, _ = cr.AddFunc("0/1 * * * * *", func() {
		fmt.Println("传统表达式: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	// 预定义表达式
	_, _ = cr.AddFunc("@every 1s", func() {
		fmt.Println("预定义表达式: ", time.Now().Format("2006-01-02 15:04:05"))
	})
	cr.Start()
	defer cr.Stop()
	select {}

}
