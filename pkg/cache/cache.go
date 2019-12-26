package cache

import (
	"dev-framework-go/conf"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/wonderivan/logger"
	"os"
	"time"
)

//import grds "github.com/go-redis/redis"
//
//func tet()  {
//	grds.NewClient()
//}
var pool *redis.Pool

func InitRedisPool() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial(conf.REDIS_NETWORK, conf.REDIS_ADDRESS)
			if err != nil {
				logger.Error("[INIT REDIS POOL ERROR]", err)
				return nil, err
			}
			_, err = c.Do("AUTH", conf.REDIS_PASS)
			if err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		// 用于检查连接被再次使用之前，连接池中的运行状况
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			fmt.Println(err)
			if err != nil {
				logger.Error("[PING REDIS ERROR]")
				return err
			}
			return nil
		},
		//池中空闲连接的最大数目
		MaxIdle: conf.REDIS_MAXIDLE,
		//池在给定时间类分配的最大连接数, 如果是0则池中的连接数没有限制
		MaxActive: conf.REDIS_MAXACTIVE,
		//在给定的间隔下，保持连接，超过这个时间后关闭连接，如果是0，则不关闭空闲连接
		IdleTimeout: conf.REDIS_IDLETIMEOUT,
		//如果wait是True并且达到了最大连接数，那么Get()方法回等待一个连接返回到池中
		Wait: conf.REDIS_WAIT,
	}
	_, err := pool.Dial()
	if err == nil {
		logger.Debug("[INIT REDIS POOL SUCCESS]")
	} else {
		os.Exit(0)
	}
}

func SetKey(key string, value interface{}, expire_time int64) bool {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		logger.Error("[SET KEY ERROR]key=%s, val=%s", key, value)
		return false
	}
	if expire_time != 0 {
		_, err := conn.Do("EXPIRE", key, expire_time)
		if err != nil {
			logger.Error("[SET KEY EXPIRE ERROR]key=%s, expire_time=%d", key, expire_time)
			return false
		}
	}
	return true
}

func GetKey(key string) interface{} {
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("GET", key))
	if err != nil {
		logger.Error("[GET KEY ERROR]key=%s", key)
		return nil
	}
	return res
}

func DelKey(key string) bool {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil {
		logger.Error("[DEL KEY ERROR]key=%s", key)
		return false
	}
	return true
}

func HSetKey(key, field, value string, expireTime int) bool {
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("HSET", key, field, value)
	if err != nil {
		logger.Error("[HSET KEY ERROR]key=%s, field=%s, value=%s", key, field, value)
		return false
	}
	if expireTime != 0 {
		_, err := conn.Do("EXPIRE", key, expireTime)
		if err != nil {
			logger.Error("[SET KEY EXPIRE ERROR]key=%s, expire_time=%d", key, expireTime)
			return false
		}
	}
	return true
}

func HGetKey(key, field string) string {
	conn := pool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("HGET", key, field))
	if err != nil {
		logger.Error("[HGET KEY ERROR]key=%s, field=%s:error:%s", key, field, err)
		return ""
	}
	return result
}

func HDelKey(key string) {
	conn := pool.Get()
	defer conn.Close()
	replay, err := redis.Strings(conn.Do("HKEYS", key))
	if err != nil {
		logger.Error("[HKEYS ERROR]key=%s", key)
		return
	}
	for k, _ := range replay {
		val := replay[k]
		_, err := conn.Do("HDEL", key, val)
		if err != nil {
			logger.Error("[HDEL KEY ERROR]key=%s", key)
		}
	}
}
