# Error loading workspace: err: exit status 1
[Error loading workspace: err: exit status 1](https://stackoverflow.com/questions/67800641/error-loading-workspace-err-exit-status-1-stderr-go-updates-to-go-sum-neede)

1. 手动删除包后，要重新开启工程；



# 包无法下载
GO MODULES 下载如果存在问题，通常是由于网络问题，首先我们可以配置代理如下：
```
GOPROXY= https://goproxy.cn,direct
```


临时配置
```
export GOPROXY=https://goproxy.cn,direct
```


永久配置
```
go env -w GOPROXY=https://goproxy.cn,direct
```


然后执行下载包到本地的命令

go mod download