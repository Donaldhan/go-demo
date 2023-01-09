#  channel
协程间的通信机制，类似共享内存模式；
![goroutine-channel](/image/goroutine-channel.jpg)


# mod 
```
go mod init godemo
```


创建包及相应的包文件，比如chanconsync


在main.go文件中应用

``` go
import (
	"godemo/chanconsync"
	"godemo/channeldemo"
	"log"
)
```

运行main文件

```
go run ./main.go
```


开启mod包管理模式，默认开启
```
go env -w GO111MODULE=on
go env -w GO111MODULE=off
```

# 附
```
go clean -modcache
```

[Go语言go mod包依赖管理工具使用详解](http://c.biancheng.net/view/5712.html)  
[分享常用的Golang工具包](https://www.jianshu.com/p/f6fa8e9b790a)   
[一个巨好用的泛型库，可以极大提高开发速度](https://github.com/samber/lo)

```
go mod download	下载依赖包到本地（默认为 GOPATH/pkg/mod 目录）
go mod edit	编辑 go.mod 文件
go mod graph	打印模块依赖图
go mod init	初始化当前文件夹，创建 go.mod 文件
go mod tidy	增加缺少的包，删除无用的包
go mod vendor	将依赖复制到 vendor 目录下
go mod verify	校验依赖
go mod why	解释为什么需要依赖
```

