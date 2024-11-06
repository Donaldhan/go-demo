package conf

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type Config struct {
	Name string
	Host string `json:",default=0.0.0.0"`
	Port int
}

var f = flag.String("f", "config.yaml", "config file")

func loadConfig() {
	flag.Parse()
	var c Config
	conf.MustLoad(*f, &c)
	// conf.MustLoad(*f, &c，conf.UseEnv()) // 额外从环境变量中加载配置
	println(c.Name)
}
