
# gateway
基于proto pb文件，提供grpc的gateway http服务；

```
protoc --descriptor_set_out=gateway/demo.pb demo.proto
```

```
curl http://localhost:8888/ping
```