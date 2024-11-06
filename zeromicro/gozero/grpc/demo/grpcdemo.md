
> 访问rpc服务
```
grpcurl -plaintext 127.0.0.1:8080 demo.Demo/Ping
```

> 列出服务
```
grpcurl -plaintext localhost:8080 list demo.Demo
grpcurl -plaintext 127.0.0.1:8080 list
```

#refer
[grpcurl](https://github.com/fullstorydev/grpcurl)  