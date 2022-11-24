package myredis

import (
	"context"
	"gs/lib/config"
	"gs/lib/mylog"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/smallnest/rpcx/log"
)

type MyRedis struct {
	client *redis.Client
}

var myRedis *MyRedis
var once sync.Once

func GetInstance() *MyRedis {
	once.Do(func() {
		myRedis = new(MyRedis)
		myRedis.init()
	})
	return myRedis
}

func (this *MyRedis) init() {
	this.client = redis.NewClient(&redis.Options{
		Addr: config.Get().RedisAddress,
	})
	_, err := this.client.Ping(context.Background()).Result()
	if err != nil {
		mylog.Error("redis connect faild " + config.Get().RedisAddress + err.Error())
		return
	}
	mylog.Debug("redis connect success " + config.Get().RedisAddress)
}

func (this *MyRedis) Del(key string) bool {
	err := this.client.Del(context.Background(), key).Err()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return true
}

func (this *MyRedis) Expire(key string, expiration time.Duration) bool {
	err := this.client.Expire(context.Background(), key, expiration).Err()
	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

//string-------------
func (this *MyRedis) Get(key string) string {
	ret, err := this.client.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.Nil {
			mylog.Error(err)
		}
	}
	return ret
}

func (this *MyRedis) Set(key string, value interface{}, expireTime time.Duration) bool {
	err := this.client.Set(context.Background(), key, value, expireTime).Err()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return true
}

func (this *MyRedis) Incr(key string) int64 {
	ret, err := this.client.Incr(context.Background(), key).Result()
	if err != nil {
		mylog.Error(err)
		return 0
	}
	return ret
}

func (this *MyRedis) SetNX(key string, value interface{}, expireTime time.Duration) bool {
	ret, err := this.client.SetNX(context.Background(), key, value, expireTime).Result()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return ret
}

func (this *MyRedis) GetLock(key string, expireTime time.Duration) bool {
	return this.SetNX(key, 1, expireTime)
}

func (this *MyRedis) ClearLock(key string) {
	this.Del(key)
}

//hash---------------
func (this *MyRedis) HGet(key, field string) string {
	ret, err := this.client.HGet(context.Background(), key, field).Result()
	if err != nil {
		mylog.Error(err)
		return ""
	}
	return ret
}

func (this *MyRedis) HSet(key, field string, value interface{}) bool {
	err := this.client.HSet(context.Background(), key, field, value).Err()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return true
}

func (this *MyRedis) HSetNX(key, field string, value interface{}) bool {
	ret, err := this.client.HSetNX(context.Background(), key, field, value).Result()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return ret
}

func (this *MyRedis) HDel(key, field string) bool {
	err := this.client.HDel(context.Background(), key, field).Err()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return true
}

func (this *MyRedis) HMSet(key string, values ...interface{}) bool {
	err := this.client.HMSet(context.Background(), key, values).Err()
	if err != nil {
		mylog.Error(err)
		return false
	}
	return true
}

func (this *MyRedis) HGetAll(key string) map[string]string {
	ret, err := this.client.HGetAll(context.Background(), key).Result()
	if err != nil {
		mylog.Error(err)
		return nil
	}
	return ret
}

func (this *MyRedis) HVals(key string) []string {
	ret, err := this.client.HVals(context.Background(), key).Result()
	if err != nil {
		mylog.Error(err)
		return nil
	}
	return ret
}

//list---------------

//set----------------

//zset---------------
