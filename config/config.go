package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

//	go get gopkg.in/yaml.v3
//
// https://gopkg.in/yaml.v3
// https://github.com/go-yaml/yaml
func ConfigDemo() {
	var c conf
	conf := c.getConf()
	log.Println("host:", conf.Host)
	log.Println("WEB3_STORAGE_TOKEN:", conf.WEB3_STORAGE_TOKEN)
	log.Println("WEB3_STORAGE_ENDPOINT:", conf.WEB3_STORAGE_ENDPOINT)
}

// 获取token
func GetWebStorageToken() string {
	var c conf
	conf := c.getConf()
	log.Println("WEB3_STORAGE_TOKEN:", conf.WEB3_STORAGE_TOKEN)
	return conf.WEB3_STORAGE_TOKEN
}

// 获取endpoint
func GetWebStorageEndPoint() string {
	var c conf
	conf := c.getConf()
	log.Println("WEB3_STORAGE_ENDPOINT:", conf.WEB3_STORAGE_ENDPOINT)
	return conf.WEB3_STORAGE_ENDPOINT
}

// profile variables
type conf struct {
	Host                  string `yaml:"host"`
	User                  string `yaml:"user"`
	Pwd                   string `yaml:"pwd"`
	Dbname                string `yaml:"dbname"`
	WEB3_STORAGE_TOKEN    string `yaml:"WEB3_STORAGE_TOKEN"`
	WEB3_STORAGE_ENDPOINT string `yaml:"WEB3_STORAGE_ENDPOINT"`
}

// 加载配置
func (c *conf) getConf() *conf {
	//得到绝对路径
	currentPath, _ := filepath.Abs("./")
	log.Println("current path:", currentPath)
	//得到他的上一级目录
	// configParent := filepath.Join(currentPath, "../")
	configParent, _ := filepath.Abs("../")
	log.Println("configParent path:", configParent)
	configPath := filepath.Join(configParent, "/conf.yaml")
	log.Println("configPath:", configPath)
	yamlFile, err := ioutil.ReadFile(configPath)
	// yamlFile, err := ioutil.ReadFile("../conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
