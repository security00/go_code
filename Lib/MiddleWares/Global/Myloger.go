package Global

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mygo/Lib/CommonFunc"
	"os"
	"runtime/debug"
	"time"
)

type customBodyLogWriter struct {
	gin.ResponseWriter

	body *bytes.Buffer
}

func (w customBodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)

	return w.ResponseWriter.Write(b)
}
func Myloger() gin.HandlerFunc {
	fmt.Println("this is record log middleware")

	// 写入文件
	src, err := os.OpenFile(CommonFunc.LogFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		fmt.Println("err:", err)
	}

	Logger := logrus.New()

	// 设置输出
	Logger.Out = src

	// 设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				logrus.Error(string(debug.Stack()))
				logrus.Error(panicErr)
			}
		}()
		bw := &customBodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}

		c.Writer = bw

		startTime := time.Now()

		data, _ := c.GetRawData() //从post中取出参数

		c.Request.Body = ioutil.NopCloser(bytes.NewReader(data)) //由于gin框架body中的数据只能取出一次，所以取出后再放回去

		c.Next()

		endTime := time.Now()
		// 执行时间
		allTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUrl := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求ip
		clientIP := c.ClientIP()

		reqDataMap := make(map[string]interface{})

		c.BindJSON(&reqDataMap)
		reqDataByte, _ := json.Marshal(reqDataMap)

		fmt.Println("===========", string(reqDataByte), "===========")

		// 换一下日期格式
		Logger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 换成json格式
		Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		// 日志变成json
		Logger.WithFields(logrus.Fields{
			"request": map[string]interface{}{
				"status_code": statusCode,
				"all_time":    allTime,
				"client_ip":   clientIP,
				"req_method":  reqMethod,
				"req_url":     reqUrl,
				"req_data":    string(data),
			},
			"response": bw.body.String(),
		}).Info()

		// 日志是字符串
		// logger.Infof("| %3d | %13v | %15s | %s | %s |",
		//  statusCode,
		//  allTime,
		//  clientIP,
		//  reqMethod,
		//  reqUrl,
		// )
	}
}
