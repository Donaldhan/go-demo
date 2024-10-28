package config

import (
	"fmt"
	"log"
	"os"
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
	log.Println("WEB3_STORAGE_ENDPOINT:", conf.WEB3_STORAGE_TOKEN)
	log.Println("jdbc:", conf.Jdbc)
	log.Println("ids:", conf.Ids)
	log.Println("Languages:", conf.Languages)
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

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
// profile variables, 注意首字母，必须可以public访问，首字母大写
type conf struct {
	Host                  string `yaml:"host"`
	WEB3_STORAGE_TOKEN    string `yaml:"WEB3_STORAGE_TOKEN"`
	WEB3_STORAGE_ENDPOINT string `yaml:"WEB3_STORAGE_ENDPOINT"`
	Jdbc                  struct {
		Driver   string `yaml:"driver"`
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
	Ids       []int    `yaml:"ids"`
	Languages []string `yaml:"languages"`
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
	//yamlFile, err := ioutil.ReadFile(configPath)
	yamlFile, err := os.ReadFile(configPath)
	// yamlFile, err := ioutil.ReadFile("../conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Println("load conf", c)
	return c
}
