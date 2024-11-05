# 安装 Google API 相关的 Protobuf 文件
```sh
go get google.golang.org/genproto
```


```sh
protoc --go_out=. --go-grpc_out=. ./micro/grpcgateway/greeter/greeter.proto
```

proto使用以下命令生成grpc stub和反向代理

```sh
protoc 
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:. \
./micro/grpcgateway/greeter/greeter.proto

```


```sh
protoc 
-I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
--go_out=plugins=grpc:. \
./micro/grpcgateway/greeter/greeter.proto

```


```sh
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:. \
./micro/grpcgateway/greeter/greeter.proto
```