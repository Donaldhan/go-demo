viper无法从远端etcd读取配置

> 环境版本

viper 版本1.19.0；
etcd client版本：3.16.0
etc 版本：3.5.7


```go
// 设置配置类型为 YAML
viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
//etcd3 需要握手
//err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", "/godemo/config.yaml")
//curl -L -k http://127.0.0.1:2379/v2/keys/godemo/config.yaml
//curl http://127.0.0.1:2379/v2/keys/godemo/config.yaml
err := viper.AddRemoteProvider("etcd", "http://127.0.0.1:2379", "/godemo/config.yaml")
if err != nil {
// Config file was found but another error was produced
log.Fatalln("error AddRemoteProvider", err)
}

if err := viper.ReadRemoteConfig(); err != nil {
	log.Fatalln("error ReadRemoteConfig", err)
}
```

# 原因分析

先来一下，如何从远程获取配置
```go
// ReadRemoteConfig attempts to get configuration from a remote source
// and read it in the remote configuration registry.

func ReadRemoteConfig() error { return v.ReadRemoteConfig() }
```


```go
package remote
...
type remoteConfigProvider struct{}

func (rc remoteConfigProvider) Get(rp viper.RemoteProvider) (io.Reader, error) {
	cm, err := getConfigManager(rp)
	if err != nil {
		return nil, err
	}
	b, err := cm.Get(rp.Path())
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
```

根据provider类型，创建对应的manager

```go
func getConfigManager(rp viper.RemoteProvider) (crypt.ConfigManager, error) {
	var cm crypt.ConfigManager
	var err error

	endpoints := strings.Split(rp.Endpoint(), ";")
	if rp.SecretKeyring() != "" {
		var kr *os.File
		kr, err = os.Open(rp.SecretKeyring())
		if err != nil {
			return nil, err
		}
		defer kr.Close()
		switch rp.Provider() {
		case "etcd":
			cm, err = crypt.NewEtcdConfigManager(endpoints, kr)
		case "etcd3":
			cm, err = crypt.NewEtcdV3ConfigManager(endpoints, kr)
		case "firestore":
			cm, err = crypt.NewFirestoreConfigManager(endpoints, kr)
		case "nats":
			cm, err = crypt.NewNatsConfigManager(endpoints, kr)
		default:
			cm, err = crypt.NewConsulConfigManager(endpoints, kr)
		}
	} else {
		switch rp.Provider() {
		case "etcd":
			cm, err = crypt.NewStandardEtcdConfigManager(endpoints)
		case "etcd3":
			cm, err = crypt.NewStandardEtcdV3ConfigManager(endpoints)
		case "firestore":
			cm, err = crypt.NewStandardFirestoreConfigManager(endpoints)
		case "nats":
			cm, err = crypt.NewStandardNatsConfigManager(endpoints)
		default:
			cm, err = crypt.NewStandardConsulConfigManager(endpoints)
		}
	}
	if err != nil {
		return nil, err
	}
	return cm, nil
}

```

针对etcd v2，
```go
func NewStandardEtcdConfigManager(machines []string) (ConfigManager, error) {
	store, err := etcd.New(machines)
	if err != nil {
		return nil, err
	}

	return NewStandardConfigManager(store)
}

```

```go
// NewStandardEtcdV3ConfigManager returns a new ConfigManager backed by etcdv3.
func NewStandardEtcdV3ConfigManager(machines []string) (ConfigManager, error) {
	store, err := etcd.NewV3(machines)
	if err != nil {
		return nil, err
	}

	return NewStandardConfigManager(store)
}

```

从Client下面可以看出，etcd(基于http的get)
```go
func (c *Client) Get(key string) ([]byte, error) {
	return c.GetWithContext(context.TODO(), key)
}

func (c *Client) GetWithContext(ctx context.Context, key string) ([]byte, error) {
	resp, err := c.keysAPI.Get(ctx, key, nil)
	if err != nil {
		return nil, err
	}
	return []byte(resp.Node.Value), nil
}

func (k *httpKeysAPI) Get(ctx context.Context, key string, opts *GetOptions) (*Response, error) {
	act := &getAction{
		Prefix: k.prefix,
		Key:    key,
	}

	if opts != nil {
		act.Recursive = opts.Recursive
		act.Sorted = opts.Sort
		act.Quorum = opts.Quorum
	}

	resp, body, err := k.client.Do(ctx, act)
	if err != nil {
		return nil, err
	}

	return unmarshalHTTPResponse(resp.StatusCode, resp.Header, body)
}

```

//再看来基于etcd3的Client
```go
func (c *Client) Get(key string) ([]byte, error) {
	return c.GetWithContext(context.TODO(), key)
}

func (c *ClientV3) Get(key string) ([]byte, error) {
	tctx, cancelFunc := context.WithTimeout(c.ctx, c.timeout)
	defer cancelFunc()
	res, err := c.keysAPI.Get(tctx, key)
	if err != nil {
		return nil, err
	}
	if res.Count != 1 {
		return nil, fmt.Errorf("getting from etcd with key [%s], res count %d not equal to 1", key, res.Count)
	}
	return res.Kvs[0].Value, nil
}
```

//kv
```go
func (kv *kv) Get(ctx context.Context, key string, opts ...OpOption) (*GetResponse, error) {
	r, err := kv.Do(ctx, OpGet(key, opts...))
	return r.get, ContextError(ctx, err)
}

func (kv *kv) Do(ctx context.Context, op Op) (OpResponse, error) {
	var err error
	switch op.t {
	case tRange:
		var resp *pb.RangeResponse
		resp, err = kv.remote.Range(ctx, op.toRangeRequest(), kv.callOpts...)
		if err == nil {
			return OpResponse{get: (*GetResponse)(resp)}, nil
		}
	case tPut:
		var resp *pb.PutResponse
		r := &pb.PutRequest{Key: op.key, Value: op.val, Lease: int64(op.leaseID), PrevKv: op.prevKV, IgnoreValue: op.ignoreValue, IgnoreLease: op.ignoreLease}
		resp, err = kv.remote.Put(ctx, r, kv.callOpts...)
		if err == nil {
			return OpResponse{put: (*PutResponse)(resp)}, nil
		}
	case tDeleteRange:
		var resp *pb.DeleteRangeResponse
		r := &pb.DeleteRangeRequest{Key: op.key, RangeEnd: op.end, PrevKv: op.prevKV}
		resp, err = kv.remote.DeleteRange(ctx, r, kv.callOpts...)
		if err == nil {
			return OpResponse{del: (*DeleteResponse)(resp)}, nil
		}
	case tTxn:
		var resp *pb.TxnResponse
		resp, err = kv.remote.Txn(ctx, op.toTxnRequest(), kv.callOpts...)
		if err == nil {
			return OpResponse{txn: (*TxnResponse)(resp)}, nil
		}
	default:
		panic("Unknown op")
	}
	return OpResponse{}, ContextError(ctx, err)
}


```

从下面可以看出，基于etcd3使用的基于kVClient：基于probuf的grpc调用
```go
type kVClient struct {
	cc *grpc.ClientConn
}

func (c *kVClient) Range(ctx context.Context, in *RangeRequest, opts ...grpc.CallOption) (*RangeResponse, error) {
out := new(RangeResponse)
err := c.cc.Invoke(ctx, "/etcdserverpb.KV/Range", in, out, opts...)
if err != nil {
return nil, err
}
return out, nil
}
```

从上面可以，最主要的问题是，etcd使用的是基于restful http模式； etcd3使用的基于probuf的grpc模式。而我们
使用的etcd restful http模式， 但etcd的版本是v3； 经测试使用etcd client/v3是可以成功获取的；那我们
知道解决方式了；




# 解决方式

## 以兼容模式运行etcd
以不安全方式运行
```shell
etcd --data-dir=/tools/etcddata --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379
```

同时provider使用etcd3
```go
// etcd3版本需要以insecure mode, 如果从V3中获取，则无法获取文件
// ```shell
// etcd --data-dir=/tools/etcddata --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379
// ```
// etcdctl put /godemo/config.yaml "$(cat config.yaml)"
// etcdctl get /godemo/config.yaml
func loadConfigRemoteEtcd3RunInsecureMode() {
// 设置配置类型为 YAML
viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
//etcd3 需要握手
err := viper.AddRemoteProvider("etcd3", "http://127.0.0.1:2379", "/godemo/config.yaml")
if err != nil {
// Config file was found but another error was produced
log.Fatalln("error AddRemoteProvider", err)
}
doRemoteConfig()
}
func doRemoteConfig() {
if err := viper.ReadRemoteConfig(); err != nil {
log.Fatalln("error ReadRemoteConfig", err)
}
// 解析到结构体
if err := viper.Unmarshal(&config); err != nil {
log.Fatalf("解析配置到结构体失败: %v", err)
}
log.Println("config:", config.Host)
s := gocron.NewScheduler()
s.Every(3).Seconds().Do(func() {
// currently, only tested with etcd support
err := viper.WatchRemoteConfig()
if err != nil {
log.Fatalln("unable to read remote config:", err)
}
// unmarshal new config into our runtime config struct. you can also use channel
// to implement a signal to notify the system of the changes
if err := viper.Unmarshal(&config); err != nil {
log.Fatalf("解析配置到结构体失败: %v", err)
}
log.Println("get host:", viper.GetString("host"))
log.Println("config:", config.Host)
})
sc := s.Start() // keep the channel
time.Sleep(60 * time.Second)
s.Clear()
fmt.Println("All task removed")
close(sc) // close the channel
<-sc      // it will happens if the channel is closed
}

```

注意如果etcd非兼容模式，provider使用etcd3， 会抛出如下错误：握手失败的错误；



## 使用etcd2版本
我的mac m3，etcd2 没有对应的包，我以docker方式运行；

同时provider使用etcd



## 使用etcd3
provider使用etcd3， 同时开启安全模式；



# 总结
etcd使用的是基于restful http模式； etcd3使用的基于probuf的grpc模式。而我们
使用的etcd restful http模式， 但etcd的版本是v3；


[etcd2 api](https://etcd.io/docs/v2.3/api/)
[etcd3 api](https://etcd.io/docs/v2.3/rfc/v3api/) 


# 附
注意：viper版本1.7.1， etcd clientV3, 存在包的冲突，grpc（1.26.0）的，难解，办法就是升级viper和etcd为最新版本；