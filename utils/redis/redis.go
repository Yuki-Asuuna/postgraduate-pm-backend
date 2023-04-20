package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"postgraduate-pm-backend/utils/helper"
	"postgraduate-pm-backend/utils/zookeeper"
	"time"
)

const (
	expire_time   = time.Hour * 12
	online_prefix = "Online_Account_"
	current_time  = "Current_Time"
)

type RedisConfig struct {
	redisAddress  string `json:"redisAddress"`
	redisPassword string `json:"redisPassword"`
	redisNetwork  string `json:"redisNetwork"`
	expireTime    int64  `json:"expireTime"`
}

var config *RedisConfig

var client *redis.Client
var ctx = context.Background()

func RedisInit() error {
	var err error
	config = &RedisConfig{}
	err = zookeeper.GetUtilsConfig("/redis", config)
	if err != nil {
		return err
	}
	client = redis.NewClient(&redis.Options{
		Addr:     config.redisAddress,
		Password: config.redisPassword,
		Network:  config.redisNetwork,
		DB:       1, // 仓库编号
	})
	return nil
}

func GetRedisClient() *redis.Client {
	return client
}

//func SetWxSessionKey(key string, openID string) error {
//	err := client.Set(ctx, key, openID, expire_time).Err()
//	if err != nil {
//		logrus.Errorf("SetWxSessionKey failed, err= %v", err)
//		return err
//	}
//	return nil
//}
//
//func GetWxOpenIDBySessionKey(sessionKey string) (string, error) {
//	val, err := client.Get(ctx, sessionKey).Result()
//	if err != nil {
//		logrus.Errorf("GetWxOpenIDBySessionKey failed, err= %v", err)
//		return "", err
//	}
//	return val, nil
//}
//
//func GetVisitorInfoBySessionKey(sessionKey string) *database.VisitorUser {
//	openID, err := GetWxOpenIDBySessionKey(sessionKey)
//	if err != nil {
//		return nil
//	}
//	user, err := database.GetVisitorUserByVisitorID(openID)
//	if err != nil {
//		return nil
//	}
//	if user == nil {
//		return nil
//	}
//	return user
//}

func GetCurrentTime() int64 {
	result, err := client.Get(ctx, current_time).Result()
	if err != nil {
		return 0
	}
	return helper.S2I64(result)
}

func SetCurrentTime(t int64) error {
	err := client.Set(ctx, current_time, t, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetOnline(userID string) error {
	err := client.Set(ctx, online_prefix+userID, 1, expire_time).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetBusy(userID string) error {
	err := client.Set(ctx, online_prefix+userID, 2, expire_time).Err()
	if err != nil {
		return err
	}
	return nil
}

func SetOffline(userID string) error {
	err := client.Del(ctx, online_prefix+userID).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetOnlineList() ([]string, error) {
	lst, err := client.Keys(ctx, online_prefix+"*").Result()
	return lst, err
}

func CheckOnline(userID string) int {
	u, err := client.Get(ctx, online_prefix+userID).Result()
	if err != nil {
		return 0
	}
	return helper.S2I(u)
}
