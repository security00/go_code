package CommonFunc

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Settings struct {
	APP_NAME string `yaml:"APP_NAME"`
	LOG_DIR  string `yaml:"LOG_DIR"`
}

var conf = Settings{}

func init() {
	env := os.Getenv("MYGOENV")
	yamlFile, err := ioutil.ReadFile("Configs/" + env + ".yaml")
	if err != nil {
		panic("get config file err:" + err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic("get config file err:" + err.Error())
	}
	result := fmt.Sprintf("%+v", conf)
	fmt.Println(result)
}

func Conf() Settings {
	return conf
}
