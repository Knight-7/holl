//Author  :     knight
//Date    :     2020/07/11 21:28:35
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :     初始化redis的连接

package dao

import (
	"holl/config"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

//RedisPool Reids的连接池
var RedisPool *redis.Pool

//InitRedis 初始化Redis连接池
func InitRedis() {
	redisConfig := config.GetRedisPoolConfig()
	RedisPool = &redis.Pool{
		//最大空闲连接数，即表示没有redis连接时依然可以保持的连接数
		MaxIdle: redisConfig.MaxIdle,
		//最大可激活的连接数，表示同时可以有的连接数
		MaxActive: redisConfig.MaxActive,
		//最大的空闲连接等待时间，超过此时间后，连接将被关闭
		IdleTimeout: time.Duration(redisConfig.IdleTimeout),
		//
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(redisConfig.Type, redisConfig.URL + ":" + redisConfig.Port)
			if err != nil {
				return nil, err
			}
			//验证redis密码
			if _, err := c.Do("AUTH", redisConfig.Auth); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
	log.Println("redis pool init success")
}

//CloseRedis 释放redis连接池
func CloseRedis() {
	log.Println("close redis pool")
	RedisPool.Close()
}