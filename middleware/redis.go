package middleware

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"godemo/tools"
)

func redisDemo() {
	c := tools.GetRedisConn() //从连接池，取一个链接
	defer c.Close()           //函数运行结束 ，把连接放回连接池
	setAndSet(c)
}

func setAndSet(c redis.Conn) {
	_, err := c.Do("Set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)

	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r1, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get book_list failed,", err)
		return
	}

	fmt.Println(r1)

	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r2, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get books failed,", err)
		return
	}

	fmt.Println(r2)
}
