module godemo

// http://c.biancheng.net/view/5712.html
// go.mod 提供了 module、require、replace 和 exclude 四个命令：
// module 语句指定包的名字（路径）；
// require 语句指定的依赖项模块；
// replace 语句可以替换依赖项模块；
// exclude 语句可以忽略依赖项模块。

// 可以使用命令go list -m -u all来检查可以升级的 package，使用go get -u need-upgrade-package升级后会将新的依赖版本更新到 go.mod * 也可以使用go get -u升级所有依赖。

// 使用 replace 替换无法直接获取的 package
// 由于某些已知的原因，并不是所有的 package 都能成功下载，比如：golang.org 下的包。

// modules 可以通过在 go.mod 文件中使用 replace 指令替换成 github 上对应的库，比如：
// replace (
//     golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
// )

// 或者
// replace golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a

// https://pkg.go.dev
require (
	// https://github.com/json-iterator/go
    // go get github.com/json-iterator/go
	github.com/json-iterator/go v1.1.12

    //  go getgithub.com/labstack/echo
	// github.com/labstack/echo v3.3.10+incompatible // indirect
	// github.com/labstack/gommon v0.3.0 // indirect
	// golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect

	// https://github.com/feng-crazy/go-utils
	// github.com/feng-crazy/go-utils v0
	// https://github.com/samber/lo
     // go get github.com/samber/lo
	github.com/samber/lo v1.37.0


    github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
    golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
)


go 1.19
