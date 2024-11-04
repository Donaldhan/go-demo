
**当前目录下的demo见，官方demo**
[grpc-go examples](https://github.com/Donaldhan/grpc-go/tree/master/examples)  


> 基于proto协议生活， go协议及go rpc文件
```sh
protoc --go-grpc_out=. --go_out=. ./middleware/rpc/gorpc/helloworld/protocol/helloworld.proto
```




参数说明

	1.	protoc：调用 Protocol Buffers 编译器 protoc，它将 .proto 文件编译成指定语言的代码文件（这里是 Go 代码）。
	2.	--go_out=.：将 Protocol Buffers 的 Go 代码生成到当前目录（.）中。此代码包含 .proto 文件中定义的消息和枚举的 Go 表示。
	3.	--go_opt=paths=source_relative：指定生成的文件路径为“源相对路径”模式。即，生成的文件相对路径与 .proto 文件路径一致，从而保持目录结构。
	•	例如，如果 .proto 文件位于 routeguide/route_guide.proto，生成的 Go 文件会在 routeguide/ 目录下。
	4.	--go-grpc_out=.：将 gRPC 服务的 Go 代码生成到当前目录中。这会生成包含 gRPC 服务接口、客户端和服务器代码的文件，便于在 Go 应用中使用 gRPC。
	5.	--go-grpc_opt=paths=source_relative：指定 gRPC 服务代码文件的路径为“源相对路径”模式，确保服务代码文件在和 .proto 文件相同的目录结构下生成。


注意事项

	•	实验性功能：--experimental_editions 是 Protocol Buffers 的实验功能，这意味着它的行为可能会在未来发生变化，建议在生产环境中慎用。
	•	兼容性：由于这是一个实验功能，使用 edition 语法的 .proto 文件可能无法在旧版本的 protoc 或没有 --experimental_editions 支持的工具中编译或运行。


```sh
protoc --experimental_editions  --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./middleware/rpc/gorpc/routeguide/protocol/route_guide.proto

```

```sh
protoc  --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./middleware/rpc/gorpc/routeguide/protocol/route_guide.proto

```

# refer
[grpc-go](https://github.com/Donaldhan/grpc-go)   
[grpc](https://grpc.io/)   
[grpc go quickstart](https://grpc.io/docs/languages/go/quickstart/)   
[grpc go quickstart cn](https://grpc.org.cn/docs/languages/go/quickstart/)  
[protobuf gotutorial](https://protobuf.com.cn/getting-started/gotutorial/)  
[grpc-go examples](https://github.com/Donaldhan/grpc-go/tree/master/examples)  