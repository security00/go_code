package CommonFunc

import (
	"path"
	"time"
)

var LogFileName string

func init()  {
	logFilePath:=Conf().LOG_DIR
	logFileName:= time.Now().Format("2006-01-02")+".log"
	//文件路径
	LogFileName = path.Join(logFilePath,logFileName)
}
