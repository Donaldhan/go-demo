
# todo

* 限流；
* 网关；
* 监控；
* 部署；




# goctl
> 
> >安装goland插件goctl

```
go install github.com/zeromicro/go-zero/tools/goctl@latest
```
**需要版本1.20以上**


```
goctl --version
goctl version 1.7.3 darwin/arm64
```

# tips

* mysql事务是基于基础的SQL来处理实现的；基础的CRUD可以使用go-zero mysql model， 涉及事务及复杂查询使用GORM更便捷；
* mongodb事务基于原生的mongodb API实现；基础的CRUD可以使用go-zero mongo model， 涉及事务及复杂查询使用MongoDB Driver更便捷；
* gozero消息队列go-queue基于kafaka和beanstalkd实现消息队列和延时队列；
* 


# refer
[website](https://go-zero.dev/)  
[go-zero](https://github.com/zeromicro/go-zero)    
[go-queue](https://github.com/Donaldhan/go-queue)  
[zero-examples](https://github.com/Donaldhan/zero-examples)    