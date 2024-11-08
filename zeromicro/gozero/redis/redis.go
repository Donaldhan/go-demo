package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func initClient() *redis.Redis {
	conf := redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		Pass:        "",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: time.Second,
	}
	return redis.MustNewRedis(conf)
}
func clientDemo() {
	rds := initClient()
	ctx := context.Background()
	//set get
	err := rds.SetCtx(ctx, "message", "hello world")
	if err != nil {
		logc.Error(ctx, err)
	}

	v, err := rds.GetCtx(ctx, "message")
	if err != nil {
		logc.Error(ctx, err)
	}
	fmt.Println("message:", v)
	//setex
	err = rds.Setex("userLock", "1", 3)
	if err != nil {
		logc.Error(ctx, err)
	}
	err = rds.Setex("userLock", "1", 3)
	if err != nil {
		log.Fatalln("setEx error", err)
	}
	//counter
	count, err := rds.Incr("counter")
	if err != nil {
		logc.Error(context.Background(), err)
	}
	log.Println("counter:", count)
	count, err = rds.Incr("counter")
	count, err = rds.Incr("counter")
	count, err = rds.Decr("counter")
	log.Println("counter:", count)
	//hset
	err = rds.Hset("user", "rain", "23")
	if err != nil {
		logc.Error(context.Background(), err)
	}
	err = rds.Hset("user", "mark", "20")
	if err != nil {
		logc.Error(context.Background(), err)
	}
	usersMap, err := rds.Hgetall("user")
	log.Println("usersMap:", usersMap)
	//list
	pushCount, err := rds.Lpush("list", "3", "4", "rain")
	log.Println("pushCount:", pushCount)
	rpopResult, err := rds.Rpop("list")
	log.Println("Rpop:", rpopResult)
	//zset
	addResutl, err := rds.Zadd("rank", 10, "rain")
	log.Println("addResutl:", addResutl)
	addResutl, err = rds.Zadd("rank", 20, "mark")
	pairs, err := rds.ZrevrangeWithScores("rank", 0, 30)
	log.Println("ZrevrangeWithScores:", pairs)

	setBitResult, err := rds.SetBit("loginStatus", 1, 1)
	log.Println("setBitResult:", setBitResult)
	setBitResult, err = rds.SetBit("loginStatus", 2, 0)
	log.Println("setBitResult:", setBitResult)
	getBitResult, err := rds.GetBit("loginStatus", 1)
	log.Println("getBitResult:", getBitResult)
}
func tryLock() {
	rds := initClient()

	lock := redis.NewRedisLock(rds, "test")

	// 设置过期时间
	lock.SetExpire(10)

	// 尝试获取锁
	acquire, err := lock.Acquire()

	switch {
	case err != nil:
		// deal err
	case acquire:
		// 获取到锁
		log.Println("tryLock acquire success")
		defer lock.Release() // 释放锁
		// 业务逻辑

	case !acquire:
		// 没有拿到锁 wait?
		log.Println("tryLock acquire fail")
	}
}
