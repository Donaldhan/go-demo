# todo
* 日志监听；
* 合约调用abi；
* 解析交易时间日志；
* 直接编码
* 改造go-web3的contract
* 多值返回改造

# usage

```
go get github.com/ethereum/go-ethereum@v1.14.11
```

# abi
```
abigen --abi ./abi/Greeter.json --pkg contract  --out ./contract/Greeter.go
abigen --abi ./abi/ComplexType.json --pkg contract  --out ./contract/base/ComplexType.go
```
Where the flags are:

--abi: Mandatory path to the contract ABI to bind to
--pkg: Mandatory Go package name to place the Go code into
--type: Optional Go type name to assign to the binding struct
--out: Optional output path for the generated Go source file (not set = stdout)

# refer
   
[go-ethereum](https://github.com/ethereum/go-ethereum)         
[geth doc](https://geth.ethereum.org/docs)  
[go-web3](https://github.com/chenzhijie/go-web3) 