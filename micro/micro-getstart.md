
# Install

```
go install github.com/micro/micro/v3@latest
```
**安装的时候，注意github.com/quic-go/quic-go依赖，目前只支持go1.21以下的，如果高于这个版本，将GOROOT换为1.21版本]**

# Running the server
```
micro server
```


Before interacting with the micro server, we need to log in with the username ‘admin’ and password ‘micro’:
```
$ micro login
```
Enter username: admin
Enter password:
Successfully logged in.


```
 micro services
 micro status
```

# Running a service
```
micro run github.com/go-micro/examples/helloworld
```



https://github.com/
# refer
[getting-started](https://micro.dev/getting-started)