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

# 检查mod冲突
1. go list -m all 命令会列出项目中所有的模块及其版本。运行此命令可以帮助查看是否存在不同模块要求相同依赖的不同版本：
```
go list -m all
```

2. go mod graph 命令会列出项目中所有依赖的模块及其依赖关系。
```
go mod graph
```

3. go mod why 命令会列出导致某个依赖被引入的模块。
```
go mod why -m  github.com/spf13/viper
```

4. 运行 go mod tidy 时，输出中会提示无效的或不必要的依赖项，帮助清理和解决版本冲突：
```
 go mod tidy
```
# 清除缓存
1. 清理所有 Go 缓存
```
go clean -modcache
```

2. 删除特定模块的缓存：
如果只想删除某个特定模块的缓存，可以手动删除 GOPATH/pkg/mod 中对应的模块文件夹。例如：
```
rm -rf $GOPATH/pkg/mod/github.com/some/module
```

3. 重建 go.mod 和 go.sum 文件：
删除缓存后，运行以下命令重新生成和清理 go.mod 和 go.sum 文件：

```
 go mod tidy
```
# go mod download 和 go get 是 Go 模块管理中两个不同的命令，各自有不同的功能和用途。以下是它们的主要区别：

1. 功能

   •	go mod download：
   •	这个命令用于下载模块依赖及其所需的所有版本到本地模块缓存中（通常在 $GOPATH/pkg/mod 下）。
   •	它不会更新 go.mod 或 go.sum 文件，仅仅是确保本地缓存中存在所需的模块。
   •	适用于你只想下载依赖而不更改依赖关系的情况。
   •	go get：
   •	这个命令用于获取模块并安装或更新模块的依赖项，同时更新 go.mod 和 go.sum 文件。
   •	如果指定了新的模块，它会将其添加到 go.mod 中；如果指定了现有模块的不同版本，它会更新该模块的版本。
   •	go get 可以用来升级或降级依赖项。


# 版本冲突
直接下载对应的版本
```
go get google.golang.org/grpc@v1.26.0
go get google.golang.org/grpc@v1.35.0
```

go get -a 命令用于安装 Go 模块及其依赖项。具体含义如下：

	•	-a（或 --all）：这个选项的主要作用是告诉 Go 工具在更新模块时也要包括该模块的所有依赖项。即使依赖项已经存在于模块缓存中，-a 选项也会强制重新下载这些依赖项。这在以下情况下非常有用：
	•	当你需要确保所有依赖项都是最新的。
	•	当某些依赖项的版本有问题，需要重新下载。


go get github.com/spf13/viper/remote@v1.7.1
go get -a github.com/spf13/viper/remote


# go install安装过程如果存在问题，则确认是否为go的版本问题
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

临时调整go版本

```
export  GOROOT=/Users/eptrusttry/Documents/go/goroot/go1.21.6
export  GOROOT=/Users/eptrusttry/Documents/go/goroot/go1.19
```

如果安装失败，则可以尝试使用brew安装grpcurl
```
brew install grpcurl
```
