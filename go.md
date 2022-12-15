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
