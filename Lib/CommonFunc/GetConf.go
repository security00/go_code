package CommonFunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
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
	User    []string `yaml:"user"`
	MQTT    MQ       `yaml:"mqtt"`
	Http    HTTP     `yaml:"http"`
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
	//fmt.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, AppConf)
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(AppConf)
	//result := fmt.Sprintf("%+v", AppConf)
	log.Fatalf("AppConf:", AppConf.Https.Port)

}

func Conf() conf {
	return *AppConf
}
