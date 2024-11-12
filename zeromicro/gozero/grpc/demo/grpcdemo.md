
> 访问rpc服务
```
grpcurl -plaintext 127.0.0.1:8080 demo.Demo/Ping
```

> 列出服务
```
grpcurl -plaintext localhost:8080 list demo.Demo
grpcurl -plaintext 127.0.0.1:8080 list
```


# 压测
```
ghz --insecure --proto=demo.proto --call=demo.Demo/Ping -d '{}' -c 90 -n 110  127.0.0.1:8080
ghz --insecure --proto=demo.proto --call=demo.Demo/Ping -d '{}' -c 100 -n 200  127.0.0.1:8080
ghz --insecure --proto=demo.proto --call=demo.Demo/Ping -d '{}' -c 200 -n 260  127.0.0.1:8080
```



#refer
[grpcurl](https://github.com/fullstorydev/grpcurl)  