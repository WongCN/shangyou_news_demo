package initialize

//var Rdb *redis.Client

//func init() {
//	initClient()
//}

// 初始化连接 (暂时未用到redis)
//func initClient() (err error) {
//	Rdb = redis.NewClient(&redis.Options{
//		Addr:     "wongcner.shop:6379",
//		Password: "password", // no password 05_set
//		DB:       0,          // use default DB
//	})
//
//	_, err = Rdb.Ping().Result()
//	if err != nil {
//		return err
//	}
//	return nil
//}
