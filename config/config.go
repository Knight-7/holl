//Author  :     knight
//Date    :     2020/07/11 21:00:11
//Version :     1.0
//Email   :     knight2347@163.com
//idea    :		读取项目的配置文件

package config

import (
	"github.com/spf13/viper"
)

//MySQLConfig MySQL的配置文件
type MySQLConfig struct {
	Name            string
	URL             string
	UserName        string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

//RedisConfig Redis连接池配置
type RedisConfig struct {
	MaxIdle     int
	MaxActive   int
	IdleTimeout int
	URL         string
	Port        string
	Type        string
	Auth        string
}

//Init 读取配置信息
func Init() error {
	//配置文件名
	fileName := "application.yaml"
	//添加配置文件路径
	viper.AddConfigPath("config")
	//指定配置配置文件
	viper.SetConfigName(fileName)
	//设置配置问价的类型
	viper.SetConfigType("yaml")

	//viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

//GetMysqlConfig 返回MySQL的配置信息
func GetMysqlConfig() *MySQLConfig {
	mysqlConfig := &MySQLConfig{
		Name:            viper.GetString("common.db.name"),
		URL:             viper.GetString("common.db.url"),
		UserName:        viper.GetString("common.db.username"),
		Password:        viper.GetString("common.db.password"),
		MaxIdleConns:    viper.GetInt("common.db.pool.maxIdleConns"),
		MaxOpenConns:    viper.GetInt("common.db.pool.maxOpenConns"),
		ConnMaxLifetime: viper.GetInt("common.db.pool.connMaxLifetime"),
	}

	return mysqlConfig
}

//GetRedisPoolConfig 返回redis的配置信息
func GetRedisPoolConfig() *RedisConfig {
	redisConfig := &RedisConfig{
		MaxIdle:     viper.GetInt("common.redis.maxIdle"),
		MaxActive:   viper.GetInt("common.redis.maxActive"),
		IdleTimeout: viper.GetInt("common.redis.idleTimeout"),
		URL:         viper.GetString("common.redis.url"),
		Port:        viper.GetString("common.redis.port"),
		Type:        viper.GetString("common.redis.type"),
		Auth:        viper.GetString("common.redis.auth"),
	}
	return redisConfig
}

//GetOssConfigFilePath 或取oss的配置文件路径
func GetOssConfigFilePath() string {
	return viper.GetString("common.oss.filePath")
}