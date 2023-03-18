package sessions

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/boj/redistore.v1"
	"postgraduate-pm-backend/database"
)

var client *redistore.RediStore

const (
	redis_address   = "124.221.197.218:6379"
	redis_password  = "ecnusyh"
	redis_network   = "tcp"
	redis_size      = 10
	redis_secretkey = "secret key"
	redis_maxage    = 12 * 60 * 60
)

func SessionInit() error {
	var err error
	client, err = redistore.NewRediStore(redis_size, redis_network, redis_address, redis_password, []byte(redis_secretkey))
	if err != nil {
		return err
	}
	client.SetMaxAge(redis_maxage)
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
