package db

import (
	"crypto/tls"
	"goSql/internal/config"
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisCon() (*redis.Options, error) {

	opt, err := redis.ParseURL(config.LoadEnv().REDIS_URL)
	if err != nil {
		log.Println("REdisEnv", config.LoadEnv().REDIS_URL)
		panic(err)
	}

	opt.TLSConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	opt.PoolSize = 500
	opt.MinIdleConns = 10
	opt.MaxIdleConns = 30

	return opt, nil
}
