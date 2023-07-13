# 引言
ipfs为分布存储协议，filecoin为ipfs的激励层，底层存储为ipfs，鼓励更多存储服务商，参与去中心化永久存储的建设中，filecoin类似以太坊的eth， filecoin的节点类似以太坊节点，并具有Ipfs的分布式存储能力；
一般不直接使用ipfs，局域网的分布式存储，可以创建类似联盟链的，ipfs cluster模式；公网的可以通过一些存储提供商将文件存储到，主网上面，比如Estuary、Web3 storage、NFT.storage。
filecoin提供了一些类似以太坊 web3 json rpc的服务，可以发送交易，gas费管理，区块等相关服务,将文件存储到filecoin，实际是向filecoin发送交易；ipfs的文件，也可以字节导入到filecoin中；


# ipfs
[ipfs](https://ipfs.tech/)  

## kudo
ipfs+filecoin的go实现；

[kubo](https://github.com/ipfs/kubo) 


## go-ipfs-api
[go-ipfs-http-client github](https://github.com/ipfs/go-ipfs-http-client)    
https://pkg.go.dev/github.com/ipfs/go-ipfs-http-client?tab=versions   
https://github.com/ipfs/go-ipfs-api  
https://docs.ipfs.tech/reference/kubo/rpc/#api-v0-object-stat   
https://sourcegraph.com/github.com/ipfs/go-ipfs-http-client/-/blob/unixfs.go   


# filecoin  

[filecoin](https://filecoin.io/zh-cn/)    
[filscan](https://filscan.io/)    
[filecoin-project](https://github.com/filecoin-project)      
[kubo-api-client](https://github.com/filecoin-project/kubo-api-client)     
[filecoin-http-api](https://github.com/filecoin-project/filecoin-http-api)    
[go-http-api](https://github.com/filecoin-project/go-http-api)    
[filecoin json rpc introduction](https://docs.filecoin.io/reference/json-rpc/introduction/)   

## lotus
lotus为filecoin实现，类似以太坊的go-ethereum等

[lotus github](https://github.com/filecoin-project/lotus)     
[lotus](https://lotus.filecoin.io/) 
[lotus filecoin](https://lotus.filecoin.io/tutorials/lotus/store-and-retrieve/set-up/) 

## test fault
[calibration](https://calibration.filscan.io/)
[calibration tutorial](https://github.com/filecoin-project/community-china/blob/master/documents/tutorial/use_cali-net_tutorial/use_cali-net_tutorial.md) 
[calibration filscan](https://calibration.filscan.io/) 
[calibration faucet](https://faucet.calibration.fildev.network) 

bafy2bzacebzhhhzhik6otsucbmrwzdgeiz5ucebwxgwnengifrb35e6mh73mg


## filecoin-http-api 

[filecoin-http-api](https://filecoin-project.github.io/filecoin-http-api/) 


# 存储服务提供商(ipfs+filecoin)
[Estuary、Web3 storage和NFT.storage](https://www.odaily.news/post/5172651)   

## estuary

[estuary](https://github.com/application-research/estuary) 
[estuary-filecoin](https://filecoin.io/zh-cn/blog/posts/estuary-filecoin/) 


## web3-storage
Web3.Storage 的内容永久存储在 Filecoin 上的存储提供商网络中，并冗余固定在 IPFS 上

ipfs的web存储服务，示例见
```
/ipfs/ipfswebstoreage.go
```
在上传文件见需要先申请token；

[web3-storage](https://web3.storage/)  
[go-w3s-client github](https://github.com/web3-storage/go-w3s-client) 
[web3.storage-filecoin](https://filecoin.io/zh-cn/blog/posts/web3.storage-filecoin/) 
http://web3.storage.ipns.localhost:8080/docs/how-tos/store/?lang=go#uploading-to-web3storage


**可选**

# 附
[IPFS和Filecoin](https://www.shilian.com/caijing/477232.html)  
[Using Filecoin with IPFS](https://filecoin.io/blog/posts/using-filecoin-with-ipfs/)    