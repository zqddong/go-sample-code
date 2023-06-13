package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

// Cache 定义适配器实现类要实现的接口
type Cache interface {
	Put(key string, value interface{})
	Get(key string) interface{}
	GetAll(keys []string) map[string]interface{}
}

// RedisCache 适配器接口
type RedisCache struct {
	conn *redis.Pool
}

// Put 缓存数据
func (rc *RedisCache) Put(key string, value interface{}) {
	if _, err := rc.conn.Get().Do("SET", key, value); err != nil {
		fmt.Println(err)
	}
}

// Get 获取缓存中指定的Key的值
func (rc *RedisCache) Get(key string) interface{} {
	value, err := redis.String(rc.conn.Get().Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

// GetAll 从缓存获取多个Key的值
func (rc *RedisCache) GetAll(keys []string) map[string]interface{} {
	intKeys := make([]interface{}, len(keys))
	for i, _ := range keys {
		intKeys[i] = keys[i]
	}

	c := rc.conn.Get()
	entries := make(map[string]interface{})
	values, err := redis.Strings(c.Do("MGET", intKeys...))
	if err != nil {
		fmt.Println(err)
		return entries
	}

	for i, k := range keys {
		entries[k] = values[i]
	}

	return entries
}

// NewRedisCache RedisCache 的工厂方法
func NewRedisCache() Cache {
	cache := &RedisCache{
		conn: &redis.Pool{
			MaxIdle:     7,
			MaxActive:   30,
			IdleTimeout: 60 * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "localhost:6379")
				if err != nil {
					fmt.Println(err)
					return nil, err
				}

				if _, err := conn.Do("SELECT", 0); err != nil {
					conn.Close()
					fmt.Println(err)
					return nil, err
				}

				return conn, nil
			},
		},
	}
	return cache
}

func main() {
	var rc Cache
	rc = NewRedisCache()
	rc.Put("Hello Go", "Nice Work")
}
