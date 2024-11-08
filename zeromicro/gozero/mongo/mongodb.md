

```shell
mongod --auth --port 27017 --dbpath=/auxiliary/mongdbdata   
```


创建管理员账号
```
use admin
```
```
db.createUser(
  {
    user: "root",
    pwd: "123456", 
    roles: [
      { role: "userAdminAnyDatabase", db: "admin" },
      { role: "readWriteAnyDatabase", db: "admin" }
    ]
  }
)
```

获取用户信息
```
db.getUsers()
```


# refer
[mongodb docs](https://www.mongodb.com/zh-cn/docs/)  
[mongodb tutorial](https://www.mongodb.com/zh-cn/docs/manual/tutorial/)  
[mongodb cn](https://www.mongodb.com/zh-cn)  