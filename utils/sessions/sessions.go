package sessions

import (
	"gopkg.in/boj/redistore.v1"
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
