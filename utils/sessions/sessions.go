package sessions

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/boj/redistore.v1"
	"postgraduate-pm-backend/database"
	"postgraduate-pm-backend/utils/zookeeper"
)

var client *redistore.RediStore

type SessionsConfig struct {
	RedisAddress   string `json:"redisAddress"`
	RedisPassword  string `json:"redisPassword"`
	RedisNetwork   string `json:"redisNetwork"`
	RedisSize      int    `json:"redisSize"`
	RedisSecretKey string `json:"redisSecretKey"`
	RedisMaxAge    int    `json:"redisMaxAge"`
}

var config *SessionsConfig

func SessionInit() error {
	var err error
	config = &SessionsConfig{}
	err = zookeeper.GetUtilsConfig("/sessions", config)
	if err != nil {
		return err
	}

	client, err = redistore.NewRediStore(config.RedisSize, config.RedisNetwork, config.RedisAddress, config.RedisPassword, []byte(config.RedisSecretKey))
	if err != nil {
		return err
	}
	client.SetMaxAge(config.RedisMaxAge)
	return nil
}

func GetSessionClient() *redistore.RediStore {
	return client
}

func GetUserIdentityNumberBySession(c *gin.Context) string {
	session, _ := client.Get(c.Request, "dotcomUser")
	ret, ok := session.Values["identityNumber"]
	if !ok {
		return ""
	}
	return ret.(string)
}

func GetUserInfoBySession(c *gin.Context) *database.User {
	currentUserIdentityNumber := GetUserIdentityNumberBySession(c)
	counsellor, err := database.GetUserByIdentityNumber(currentUserIdentityNumber)
	if err != nil {
		return nil
	}
	return counsellor
}
