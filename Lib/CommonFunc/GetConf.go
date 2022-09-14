package CommonFunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type HTTP struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type MQ struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type conf struct {
	LOGPATH string   `yaml:"log-path"`
	User    []string `yaml:"user"`
	MQTT    MQ       `yaml:"mqtt"`
	Http    HTTP     `yaml:"http"`
}

type Config struct {
	conf conf `yaml:"config"`
}

var AppConf Config

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
	result := fmt.Sprintf("%+v", &AppConf)
	log.Fatalf("AppConf:", result)

	//yamlFile, err := os.Open("Configs/" + env + ".yaml")
	//if err != nil {
	//	panic("get config file err:" + err.Error())
	//}
	//yaml.NewDecoder(yamlFile).Decode(AppConf)
	//log.Fatalf("conf: ", AppConf)
}

func Conf() conf {
	return AppConf.conf
}
