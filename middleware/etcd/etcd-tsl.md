> todo当前验证存在问题，待验证

签发证书用于为 etcd 等服务启用安全的 TLS 加密通信。在生产环境中，建议通过可靠的证书颁发机构 (CA) 或内部的 CA 来签发证书，以确保通信的安全性。以下是使用 openssl 来创建自签名证书和签发证书的详细步骤：

# 创建自签名 CA 证书

首先，生成一个自签名 CA 证书，用于签发服务器和客户端的证书。

```
# 创建 CA 私钥
openssl genpkey -algorithm RSA -out ca-key.pem -pkeyopt rsa_keygen_bits:2048

# 使用 CA 私钥创建自签名的 CA 证书
openssl req -new -x509 -key ca-key.pem -out ca-cert.pem -days 365 -subj "/CN=etcd-ca"
```


# 为 etcd 服务器签发证书

使用 CA 证书签发 etcd 服务器证书，并生成服务器私钥。

	1.	生成服务器私钥：
```
openssl genpkey -algorithm RSA -out server-key.pem -pkeyopt rsa_keygen_bits:2048

```

	2.	生成服务器证书签名请求 (CSR)：

```
openssl req -new -key server-key.pem -out server.csr -subj "/CN=etcd-server"
```

	3.	使用 CA 证书签发服务器证书：

```
openssl x509 -req -in server.csr -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -days 365
```

生成的 server-cert.pem 和 server-key.pem 文件即为服务器的证书和密钥，ca-cert.pem 是 CA 证书。

# 为 etcd 客户端签发证书

如果客户端也需要加密访问 etcd，则可以为客户端生成单独的证书。

	1.	生成客户端私钥：
```
openssl genpkey -algorithm RSA -out client-key.pem -pkeyopt rsa_keygen_bits:2048
```

    2.	生成客户端证书签名请求 (CSR)：
```
openssl req -new -key client-key.pem -out client.csr -subj "/CN=etcd-client"  
```

    3.	使用 CA 证书签发客户端证书：
```
openssl x509 -req -in server.csr -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -days 365
```

# 配置 etcd 使用证书

将生成的证书文件 (server-cert.pem、server-key.pem、ca-cert.pem) 配置到 etcd 服务中：
```
etcd \
  --cert-file=/tools/etcddata/server-cert.pem \
  --key-file=/tools/etcddata/server-key.pem \
  --trusted-ca-file=/tools/etcddata/ca-cert.pem \
  --client-cert-auth \
  --peer-cert-file=/tools/etcddata/server-cert.pem \
  --peer-key-file=/tools/etcddata/server-key.pem \
  --peer-trusted-ca-file=/tools/etcddata/ca-cert.pem \
  --listen-client-urls=https://127.0.0.1:2379 \
  --advertise-client-urls=https://127.0.0.1:2379 \
  --listen-peer-urls=https://127.0.0.1:2380 \
  --initial-advertise-peer-urls=https://127.0.0.1:2380
```

# 验证证书配置

启动 etcd 后，客户端可以使用 etcdctl 或其他工具通过证书进行访问。示例：

```
etcdctl  --cacert --endpoints=https://127.0.0.1:2379 \
--cacert=/tools/etcddata/ca-cert.pem \
--cert=/tools/etcddata/client-cert.pem \
--key=/tools/etcddata/client-key.pem \
endpoint health
```

注意事项

	•	证书有效期：建议定期更新证书以确保安全性。
	•	证书管理：在生产环境中，可以使用证书管理工具来简化证书签发和轮换流程。
注意事项

	1.	确保客户端与服务器证书的 CA 是相同的，才能顺利通过 TLS 验证。
	2.	如果 etcd 使用了自签名证书，客户端连接必须使用 --cacert 参数指定该证书。

# put & get 操作
```
etcdctl  --cacert --endpoints=https://127.0.0.1:2379 \
        --cacert=/tools/etcddata/ca-cert.pem \
        --cert=/tools/etcddata/client-cert.pem \
        --key=/tools/etcddata/client-key.pem \
        put  /godemo/config.yaml "$(cat config.yaml)"  
```

```
etcdctl --endpoints=https://127.0.0.1:2379 \
        --cacert=/tools/etcddata/ca-cert.pem \
        --cert=/tools/etcddata/client-cert.pem \
        --key=/tools/etcddata/client-key.pem \
        get /godemo/config.yaml

```


# refer
[Transport security model](https://etcd.io/docs/v3.5/op-guide/security/)  