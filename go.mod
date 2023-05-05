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
// go mod tidy ，下载更新依赖
// go install这种下载依赖的方式其实是通过go get的方式去下载的
// go install -x 加上-x命令，可以查看更多的错误信息
// go mod download
// go mod graph
require (
	//https://github.com/Donaldhan/martini
	// go get github.com/Donaldhan/martini
	github.com/Donaldhan/martini v1.0.0

	// https://github.com/go-martini/martini
	// go get github.com/go-martini/martini
	// go mod download github.com/codegangsta/inject
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0
	// https://pkg.go.dev/github.com/ipfs/go-ipfs-http-client?tab=versions
	// https://github.com/ipfs/go-ipfs-http-client
	github.com/ipfs/go-ipfs-http-client v0.4.0

	// https://github.com/json-iterator/go
	// go get github.com/json-iterator/go
	github.com/json-iterator/go v1.1.12
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect

	// https://github.com/pkg/profile
	// go get github.com/pkg/profile
	github.com/pkg/profile v1.7.0

	//  go getgithub.com/labstack/echo
	// github.com/labstack/echo v3.3.10+incompatible // indirect
	// github.com/labstack/gommon v0.3.0 // indirect
	// golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413 // indirect

	// https://github.com/feng-crazy/go-utils
	// github.com/feng-crazy/go-utils v0
	// https://github.com/samber/lo
	// go get github.com/samber/lo
	github.com/samber/lo v1.37.0
	golang.org/x/exp v0.0.0-20220303212507-bbda1eaf7a17 // indirect
)

require (
	github.com/ipfs/go-ipfs-files v0.1.1
	github.com/ipfs/interface-go-ipfs-core v0.7.0 // indirect
	github.com/multiformats/go-multiaddr v0.5.0
)

require (
	github.com/btcsuite/btcd v0.21.0-beta // indirect
	github.com/crackcomm/go-gitignore v0.0.0-20170627025303-887ab5e44cc3 // indirect
	github.com/felixge/fgprof v0.9.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/pprof v0.0.0-20211214055906-6f57359322fd // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/ipfs/bbloom v0.0.4 // indirect
	github.com/ipfs/go-block-format v0.0.3 // indirect
	github.com/ipfs/go-blockservice v0.3.0 // indirect
	github.com/ipfs/go-cid v0.1.0 // indirect
	github.com/ipfs/go-datastore v0.5.0 // indirect
	github.com/ipfs/go-ipfs-blockstore v1.2.0 // indirect
	github.com/ipfs/go-ipfs-cmds v0.7.0 // indirect
	github.com/ipfs/go-ipfs-ds-help v1.1.0 // indirect
	github.com/ipfs/go-ipfs-exchange-interface v0.1.0 // indirect
	github.com/ipfs/go-ipfs-util v0.0.2 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.5 // indirect
	github.com/ipfs/go-ipld-format v0.4.0 // indirect
	github.com/ipfs/go-ipld-legacy v0.1.0 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.3.0 // indirect
	github.com/ipfs/go-merkledag v0.6.0 // indirect
	github.com/ipfs/go-metrics-interface v0.0.1 // indirect
	github.com/ipfs/go-path v0.3.0 // indirect
	github.com/ipfs/go-unixfs v0.3.1 // indirect
	github.com/ipfs/go-verifcid v0.0.1 // indirect
	github.com/ipld/go-codec-dagpb v1.3.0 // indirect
	github.com/ipld/go-ipld-prime v0.11.0 // indirect
	github.com/jbenet/goprocess v0.1.4 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/libp2p/go-buffer-pool v0.0.2 // indirect
	github.com/libp2p/go-libp2p-core v0.8.6 // indirect
	github.com/libp2p/go-openssl v0.0.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.0.3 // indirect
	github.com/multiformats/go-base36 v0.1.0 // indirect
	github.com/multiformats/go-multibase v0.0.3 // indirect
	github.com/multiformats/go-multicodec v0.4.1 // indirect
	github.com/multiformats/go-multihash v0.1.0 // indirect
	github.com/multiformats/go-varint v0.0.6 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/polydawn/refmt v0.0.0-20201211092308-30ac6d18308e // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/spacemonkeygo/spacelog v0.0.0-20180420211403-2296661a0572 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20200123233031-1cdf64d27158 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	lukechampine.com/blake3 v1.1.6 // indirect
)

go 1.19
