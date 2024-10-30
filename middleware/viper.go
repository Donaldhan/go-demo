package middleware

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jasonlvhit/gocron"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Host                  string `mapstructure:"host"`
	WEB3_STORAGE_TOKEN    string `mapstructure:"WEB3_STORAGE_TOKEN"`
	WEB3_STORAGE_ENDPOINT string `mapstructure:"WEB3_STORAGE_ENDPOINT"`
	Jdbc                  struct {
		Driver   string `mapstructure:"driver"`
		Url      string `mapstructure:"url"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	}
	Ids       []int    `mapstructure:"ids"`
	Languages []string `mapstructure:"languages"`
	Reload    bool     `mapstructure:"reload"`
	Times     int      `mapstructure:"times"`
}

var config Config

func watchConfig() {
	loadConfig()
	log.Println("host:", config.Host)
	s := gocron.NewScheduler()
	s.Every(3).Seconds().Do(func() {
		//动态监听，必须如下方式，才能拿到最新的值
		log.Println("get host:", viper.GetString("host"))
	})
	sc := s.Start() // keep the channel
	time.Sleep(30 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc) // close the channel
	<-sc      // it will happens if the channel is closed
}
func viperConfig() {
	loadConfig()
	log.Println("host:", config.Host)
	log.Println("jdbc:", config.Jdbc)
	log.Println("ids:", config.Ids)
	log.Println("languages:", config.Languages)
	log.Println("reload:", config.Reload)
	log.Println("times:", config.Times)
}
func viperDemo() {
	loadConfig()
	log.Println("host:", viper.GetString("host"))
	log.Println("jdbc:", viper.GetStringMapString("jdbc"))
	log.Println("ids:", viper.GetIntSlice("ids"))
	log.Println("languages:", viper.GetStringSlice("languages"))
	log.Println("reload:", viper.GetBool("reload"))
	log.Println("times:", viper.GetInt("times"))
}
func loadConfig() {
	//绑定环境变量
	//export MY_APP_NAME=MyEnvironmentApp
	//export APP_PORT=9090

	// 绑定单个环境变量
	viper.BindEnv("APP_NAME", "MY_APP_NAME")

	// 读取环境变量并提供默认值
	viper.SetDefault("APP_PORT", 8080)
	viper.BindEnv("APP_PORT") // 绑定到 APP_PORT 环境变量
	// 打印测试
	fmt.Println("App Name:", viper.GetString("APP_NAME"))
	fmt.Println("App Port:", viper.GetInt("APP_PORT"))

	viper.SetDefault("PROFILE", "dev")
	viper.BindEnv("PROFILE")

	profile := viper.GetString("PROFILE")
	log.Println("profile:", profile)
	fileName := "config-" + profile
	log.Println("profile fileName:", fileName)
	viper.SetConfigName(fileName) // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	//viper.AddConfigPath("/etc/appname/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	viper.AddConfigPath("..")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatalln("error reading config file", err)
		}
	}
	// 解析到结构体
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置到结构体失败: %v", err)
	}

	// 监听文件变化
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name, e.Op)
		if e.Op == fsnotify.Write || e.Op == fsnotify.Create {
			log.Printf("配置文件 %s 已更改，重新加载...", e.Name)
			if err := viper.MergeInConfig(); err != nil {
				log.Printf("重新加载配置失败: %v", err)
			}
		}
	})
	viper.WatchConfig()

}

// 命令行参数
// ./myapp --config=config.yaml --port=9090
// go run main.go --config=config.yaml --port=9090
func InitFlags() {
	// 定义命令行标志
	pflag.String("config", "", "配置文件路径")
	pflag.Int("port", 8080, "服务端口号")

	// 解析命令行标志
	pflag.Parse()

	// 将 Pflags 绑定到 Viper
	viper.BindPFlags(pflag.CommandLine)

	// 打印测试
	fmt.Println("Config File:", viper.GetString("config"))
	fmt.Println("Port:", viper.GetInt("port"))
}
