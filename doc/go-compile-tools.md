# go build

go build 编译时的附加参数:

- v	编译时显示包名
-p n	开启并发编译，默认情况下该值为 CPU 逻辑核数
-a	强制重新构建
-n	打印编译时会用到的所有命令，但不真正执行
-x	打印编译时会用到的所有命令
-race	开启竞态检测 


```
go build -o main -v -x
```

# go clean
Go语言中go clean命令可以移除当前源码包和关联源码包里面编译生成的文件，这些文件包括以下几种：
执行go build命令时在当前目录下生成的与包名或者 Go 源码文件同名的可执行文件。在 Windows 下，则是与包名或者 Go 源码文件同名且带有“.exe”后缀的文件。
执行go test命令并加入-c标记时在当前目录下生成的以包名加“.test”后缀为名的文件。在 Windows 下，则是以包名加“.test.exe”后缀的文件。
执行go install命令安装当前代码包时产生的结果文件。如果当前代码包中只包含库源码文件，则结果文件指的就是在工作区 pkg 目录下相应的归档文件。如果当前代码包中只包含一个命令源码文件，则结果文件指的就是在工作区 bin 目录下的可执行文件。
在编译 Go 或 C 源码文件时遗留在相应目录中的文件或目录 。包括：“_obj”和“_test”目录，名称为“_testmain.go”、“test.out”、“build.out”或“a.out”的文件，名称以“.5”、“.6”、“.8”、“.a”、“.o”或“.so”为后缀的文件。这些目录和文件是在执行go build命令时生成在临时目录中的。

go clean命令就像 Java 中的maven clean命令一样，会清除掉编译过程中产生的一些文件。在 Java 中通常是 .class 文件，而在Go语言中通常是上面我们所列举的那些文件。

```
go clean -i -n -x
```

通过上面的示例可以看出，go clean命令还可以指定一些参数。对应的参数的含义如下所示：
-i 清除关联的安装的包和可运行文件，也就是通过go install安装的文件；
-n 把需要执行的清除命令打印出来，但是不执行，这样就可以很容易的知道底层是如何运行的；
-r 循环的清除在 import 中引入的包；
-x 打印出来执行的详细命令，其实就是 -n 打印的执行版本；
-cache 删除所有go build命令的缓存
-testcache 删除当前包所有的测试结果


# go run

go run不会在运行目录下生成任何文件，可执行文件被放在临时文件中被执行，工作目录被设置为当前目录。在go run的后部可以添加参数，这部分参数会作为代码可以接受的命令行输入提供给程序。

```
 go run main.go
```

# go fmt
gofmt 介绍
Go语言的开发团队制定了统一的官方代码风格，并且推出了 gofmt 工具（gofmt 或 go fmt）来帮助开发者格式化他们的代码到统一的风格。

gofmt 是一个 cli 程序，会优先读取标准输入，如果传入了文件路径的话，会格式化这个文件，如果传入一个目录，会格式化目录中所有 .go 文件，如果不传参数，会格式化当前目录下的所有 .go 文件。

gofmt 默认不对代码进行简化，使用-s参数可以开启简化代码功能

# go install

go install 命令的功能和前面一节《go build命令》中介绍的 go build 命令类似，附加参数绝大多数都可以与 go build 通用。go install 只是将编译的中间文件放在 GOPATH 的 pkg 目录下，以及固定地将编译结果放在 GOPATH 的 bin 目录下。

这个命令在内部实际上分成了两步操作：第一步是生成结果文件（可执行文件或者 .a 包），第二步会把编译好的结果移到 $GOPATH/pkg 或者 $GOPATH/bin。

```
go install
```

go install 的编译过程有如下规律：
go install 是建立在 GOPATH 上的，无法在独立的目录里使用 go install。
GOPATH 下的 bin 目录放置的是使用 go install 生成的可执行文件，可执行文件的名称来自于编译时的包名。
go install 输出目录始终为 GOPATH 下的 bin 目录，无法使用-o附加参数进行自定义。
GOPATH 下的 pkg 目录放置的是编译期间的中间文件。


# go get

go get 命令可以借助代码管理工具通过远程拉取或更新代码包及其依赖包，并自动完成编译和安装。整个过程就像安装一个 App 一样简单。

这个命令可以动态获取远程代码包，目前支持的有 BitBucket、GitHub、Google Code 和 Launchpad。在使用 go get 命令前，需要安装与远程包匹配的代码管理工具，如 Git、SVN、HG 等，参数中需要提供一个包名。

这个命令在内部实际上分成了两步操作：第一步是下载源码包，第二步是执行 go install。下载源码包的 go 工具会自动根据不同的域名调用不同的源码工具，对应关系如下：
BitBucket (Mercurial Git)
GitHub (Git)
Google Code Project Hosting (Git, Mercurial, Subversion)
Launchpad (Bazaar)

所以为了 go get 命令能正常工作，你必须确保安装了合适的源码管理工具，并同时把这些命令加入你的 PATH 中。其实 go get 支持自定义域名的功能。

参数介绍：
-d 只下载不安装
-f 只有在你包含了 -u 参数的时候才有效，不让 -u 去验证 import 中的每一个都已经获取了，这对于本地 fork 的包特别有用
-fix 在获取源码之后先运行 fix，然后再去做其他的事情
-t 同时也下载需要为运行测试所需要的包
-u 强制使用网络去更新包和它的依赖包
-v 显示执行的命令


go get+ 远程包
默认情况下，go get 可以直接使用。例如，想获取 go 的源码并编译，使用下面的命令行即可：


获取前，请确保 GOPATH 已经设置。Go 1.8 版本之后，GOPATH 默认在用户目录的 go 文件夹下。

cellnet 只是一个网络库，并没有可执行文件，因此在 go get 操作成功后 GOPATH 下的 bin 目录下不会有任何编译好的二进制文件。

```
$ go get github.com/Donaldhan/martini 
```
# go generate

go generate命令是在Go语言 1.4 版本里面新添加的一个命令，当运行该命令时，它将扫描与当前包相关的源代码文件，找出所有包含//go:generate的特殊注释，提取并执行该特殊注释后面的命令

# go test
单元测试——测试和验证代码的框架
要开始一个单元测试，需要准备一个 go 源码文件，在命名文件时需要让文件必须以_test结尾。默认的情况下，go test命令不需要任何的参数，它会自动把你源码包下面所有 test 文件测试完毕，当然你也可以带上参数。

这里介绍几个常用的参数：
-bench regexp 执行相应的 benchmarks，例如 -bench=.；
-cover 开启测试覆盖率；
-run regexp 只运行 regexp 匹配的函数，例如 -run=Array 那么就执行包含有 Array 开头的函数；
-v 显示测试的详细命令。

```
go test -v
```
[go test](http://c.biancheng.net/view/124.html) 


# go pprof
Go语言工具链中的 go pprof 可以帮助开发者快速分析及定位各种性能问题，如 CPU 消耗、内存分配及阻塞分析。

性能分析首先需要使用 runtime.pprof 包嵌入到待分析程序的入口和结束处。runtime.pprof 包在运行时对程序进行每秒 100 次的采样，最少采样 1 秒。然后将生成的数据输出，让开发者写入文件或者其他媒介上进行分析。

go pprof 工具链配合 Graphviz 图形化工具可以将 runtime.pprof 包生成的数据转换为 PDF 格式，以图片的方式展示程序的性能分析结果

```
$ go build -o cpu cpu.go
$ ./cpu
$ go tool pprof --pdf cpu cpu.pprof > cpu.pdf
```