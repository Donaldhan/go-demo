# etcd
https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/go%E6%93%8D%E4%BD%9Cetcd/etcd%E4%BB%8B%E7%BB%8D.html

[etcd](https://etcd.io/)       
[etcd](https://github.com/etcd-io/etcd)   

```shell
etcdctl put name rain
etcdctl get name
etcdctl del name
```

[//]: # (todo upload file)

etcdctl put /godemo/config.yaml "$(cat config.yaml)"  
etcdctl get /godemo/config.yaml


```
go get -u go.etcd.io/etcd/client/v3
```

以不安全方式运行
```shell
etcd --data-dir=/tools/etcddata --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379
```
# etcd2 docker

在etcd github仓库下， 编译， 然后构建镜像：
```
docker build -t etcd:v2.3.8-arm64 .
```

```
docker run -d --name etcd2-arm \
-p 2379:2379 \
-p 2380:2380 \
etcd:v2.3.8-arm64 \
etcd --listen-client-urls http://0.0.0.0:2379 \
--advertise-client-urls http://0.0.0.0:2379
```

# etcd v3 tsl模式



# etcd 图形化管理工具
[etcdmanager](https://github.com/gtamas/etcdmanager)  
[etcd-browser](https://github.com/henszey/etcd-browser)  
[boot4go-etcdv3-browser](https://github.com/gohutool/boot4go-etcdv3-browser)
[redisant](http://www.redisant.cn/etcd)
[k8s-etcd-ui](https://github.com/comqx/k8s-etcd-ui)    
dockers etcd2 keeper v3dockers


## etcd-browser
需要构建镜像，在容器里跑, 不太好用
```
docker run --rm --name etcd-browser -p 0.0.0.0:8000:8000 --env ETCD_HOST=127.0.0.1  --env ETCD_PORT=2379
```