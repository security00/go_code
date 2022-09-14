package CommonFunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type HTTPS struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}
type MQ struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type conf struct {
	Mysql   []string `yaml:"mysql"`
	MQTT    MQ       `yaml:"mqtt"`
	Https   HTTPS    `yaml:"https"`
	LOGPATH string   `yaml:"log-path"`
}

var AppConf = new(conf)

func init() {
	env := os.Getenv("MYGOENV")
	yamlFile, err := ioutil.ReadFile("Configs/" + env + ".yaml")
	if err != nil {
		panic("get config file err:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, AppConf)
	if err != nil {
		fmt.Println(err.Error())
	}
	//result := fmt.Sprintf("%+v", AppConf)
	//log.Fatalf("AppConf:", AppConf.Mysql)

}

func Conf() conf {
	return *AppConf
}
