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
goctl model mysql datasource -url="username:password@tcp(127.0.0.1:3306)/dbname" -table="table_name" -dir="./model"
```

# 生成mongo模型
```
goctl model mongo --type user --dir .
```

# 格式化api文档
```
goctl api format --dir demo.api
```

# 生成单体服务
```
goctl quickstart --service-type mono
```
与api demo类似