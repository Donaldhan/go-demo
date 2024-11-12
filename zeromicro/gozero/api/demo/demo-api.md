```
$ curl --request GET 'http://127.0.0.1:8888/from/me'
```

```
curl --request GET 'http://127.0.0.1:8888/from/me1'
```

# 压测
检查限流策略：并发，超时
```
hey -z 1s -c 200 -q 1 'http://127.0.0.1:8888/from/me1'
```

# 健康检查
```
 curl -i http://127.0.0.1:6060/healthz
```