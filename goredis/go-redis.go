package goredis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var Rdb redis.Conn

func init() {
	Rdb = GetRedisPool().Get()
}

type RedisConfig struct {
	RedisConn      string
	RedisDbNum     int
	RedisPassword  string
	RedisMaxIdle   int
	RedisMaxActive int
}

func NewRedisConfig() *RedisConfig {
	db := &RedisConfig{RedisDbNum: 0, RedisMaxIdle: 1, RedisMaxActive: 10}
	db.RedisConn = "127.0.0.1:6379"
	return db
}

func GetRedisPool() *redis.Pool {
	config := NewRedisConfig()
	//连接地址
	RedisConn := config.RedisConn
	//db分区
	RedisDbNum := config.RedisDbNum
	//密码
	RedisPassword := config.RedisPassword

	//建立连接池
	return &redis.Pool{
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: config.RedisMaxIdle,
		//最大的激活连接数，表示同时最多有N个连接
		MaxActive: config.RedisMaxActive,
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 300 * time.Second,
		//建立连接
		Dial: func() (redis.Conn, error) {
			logError(RedisConn)
			c, err := redis.Dial("tcp", RedisConn)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			if RedisPassword != "" {
				if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
					return nil, fmt.Errorf("redis auth password error: %s", authErr)
				}
			}
			//选择分区
			c.Do("SELECT", RedisDbNum)
			return c, nil
		},
		//ping
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}

}

func SetString(key, v string) (interface{}, error) {
	res, err := Rdb.Do("SET", key, v)
	if err != nil {
		logError("set error", err.Error())
		return false, err
	}
	b := false
	if res.(string) == "OK" {
		b = true
	}
	return b, nil
}

func GetString(key string) (string, error) {
	val, err := redis.String(Rdb.Do("GET", key))
	if err != nil {
		logError("get error", err.Error())
		return "", err
	}
	return val, nil
}


func Expire(key string, ex int) error {
	_, err := Rdb.Do("EXPIRE", key, ex)
	if err != nil {
		logError("set error", err.Error())
		return err
	}
	return nil
}

func SetExString(key, v string, ex int) error {
	_, err := Rdb.Do("SETEX", key, ex, v)
	if err != nil {
		logError("set error", err.Error())
		return err
	}
	return nil
}


func Exists(key string) bool {
	b, err := redis.Bool(Rdb.Do("EXISTS", key))
	if err != nil {
		logError("Exists error:", err)
		return false
	}
	return b
}


func Del(key string) error {
	_, err := Rdb.Do("DEL", key)
	if err != nil {
		logError("del error:", err)
		return err
	}
	return nil
}

// data 可以是map 或者结构体，最后都会转换成json
func SetJson(key string, data interface{}) error {
	value, _ := json.Marshal(data)
	n, _ := Rdb.Do("setNx", key, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}


func GetJson(key string) (map[string]interface{}, error) {
	var jsonData map[string]interface{}
	bv, err := redis.Bytes(Rdb.Do("GET", key))
	if err != nil {
		logError("get json error", err.Error())
		return nil, err
	}
	errJson := json.Unmarshal(bv, &jsonData)
	if errJson != nil {
		logError("json nil", errJson.Error())
		return nil, err
	}
	return jsonData, nil
}


func HSet(key string, field string, data interface{}) error {
	_, err := Rdb.Do("hSet", key, field, data)
	if err != nil {
		logError("hSet error", err.Error())
		return err
	}
	return nil
}

func HGet(key, field string) (interface{}, error) {
	data, err := Rdb.Do("hGet", key, field)
	if err != nil {
		logError("hGet error", err.Error())
		return nil, err
	}
	return data, nil
}

// 获取所有hash
func HGetAll(key string) (map[string]string, error) {
	data, err2 := redis.StringMap(Rdb.Do("HGetAll", key))
	_, err := data, err2
	if err != nil {
		logError("hGetAll error", err.Error())
		return nil, err
	}
	return data, nil
}
//自增1
func Incr(key string) error {
	_, err := Rdb.Do("INCR", key)
	if err != nil {
		logError("INCR error", err.Error())
		return err
	}
	return nil

}
//自增n
func IncrBy(key string, n int) error {
	_, err := Rdb.Do("IncrBy", key, n)
	if err != nil {
		logError("IncrBy error", err.Error())
		return err
	}
	return nil
}
// 自减1
func Decr(key string) error {
	_, err := Rdb.Do("DECR", key)
	if err != nil {
		logError("DECR error", err.Error())
		return err
	}
	return nil
}

// 自减n
func DecrBy(key string, n int) error {
	_, err := Rdb.Do("DecrBy", key, n)
	if err != nil {
		logError("DecrBy error", err.Error())
		return err
	}
	return nil
}

// 增加集合
func SAdd(key, v string) error {
	_, err := Rdb.Do("SAdd", key, v)
	if err != nil {
		logError("sAdd error", err.Error())
		return err
	}
	return nil
}

// 获取所有集合列表
func SMembers(key string) (interface{}, error) {
	data, err := redis.Strings(Rdb.Do("SMembers", key))
	if err != nil {
		logError("json nil", err)
		return nil, err
	}
	return data, nil
}

//判断 member 元素是否集合 key 的成员。
func SISMembers(key, v string) bool {
	b, err := redis.Bool(Rdb.Do("SISMembers", key, v))
	if err != nil {
		logError("SISMembers error", err.Error())
		return false
	}
	return b
}

// 打印日志
func logError(str ...interface{}) {
	fmt.Println(str...)
}

