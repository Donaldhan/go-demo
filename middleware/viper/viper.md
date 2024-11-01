

Viper支持多种配置源，并处理了它们之间的优先级问题，按优先级从高到低的顺序，层级如下：

* 命令行参数
* 环境变量
* 本地配置文件
* 远程配置文件
* 默认值


可以读取远端文件，并监听文件变化

监听方式
```
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
```

注意： 动态监听，必须如下方式，才能拿到最新的值
```
log.Println("get host:", viper.GetString("host"))
```


# refer
[viper](https://github.com/spf13/viper)  
[浅谈Golang配置管理](https://juejin.cn/post/7246304095375622203)    
[golang升级etcd解决 grpc兼容性问题](https://l1905.github.io/golang/2021/09/06/golang-etcd-grpc-incompatible/) 