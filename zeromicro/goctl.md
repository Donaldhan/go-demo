# 生成apidemo
执行指令生成 demo 服务
```
goctl api new demo
```

```
go mod init apidemo
```

```
go mod tidy
```


# 生成grpc demo
```
goctl rpc new demo
```


# 根据sql生成模型
```
goctl model mysql ddl --src user.sql --dir .
```

# 
```
goctl model mongo --type user --dir .
```