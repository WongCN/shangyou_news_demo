package initialize

import "github.com/go-redis/redis"

var Rdb *redis.Client

func init() {
	initClient()
}

//初始化连接
func initClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "wongcner.shop:6379",
		Password: "password", // no password 05_set
		DB:       0,          // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
