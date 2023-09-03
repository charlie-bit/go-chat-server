package redis

import (
	"chat_socket/server/config"
	"errors"
	"fmt"
	"github.com/charlie-bit/utils/third_party/go-redis"
)

func NewRedis() (redis.UniversalClient, error) {
	if len(config.Cfg.Redis.Address) == 0 {
		return nil, errors.New("redis address is empty")
	}
	var rdb redis.UniversalClient
	if len(config.Cfg.Redis.Address) > 1 {
		rdb = redis.NewClusterClient(
			&redis.ClusterOptions{
				Addrs:      config.Cfg.Redis.Address,
				Username:   config.Cfg.Redis.Username,
				Password:   config.Cfg.Redis.Password, // no password set
				PoolSize:   50,
				MaxRetries: 3,
			},
		)
	} else {
		rdb = redis.NewClient(
			&redis.Options{
				Addr:       config.Cfg.Redis.Address[0],
				Username:   config.Cfg.Redis.Username,
				Password:   config.Cfg.Redis.Password, // no password set
				DB:         0,                         // use default DB
				PoolSize:   100,                       // connection pool size
				MaxRetries: 3,
			},
		)
	}

	var err error = nil
	err = rdb.Ping().Err()
	if err != nil {
		return nil, fmt.Errorf("redis ping %w", err)
	}
	return rdb, err
}
