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


# etcd 图形化管理工具
[etcdmanager](https://github.com/gtamas/etcdmanager) 
[boot4go-etcdv3-browser](https://github.com/gohutool/boot4go-etcdv3-browser)
[redisant](http://www.redisant.cn/etcd)
[k8s-etcd-ui](https://github.com/comqx/k8s-etcd-ui)    
dockers etcd2 keeper v3dockers