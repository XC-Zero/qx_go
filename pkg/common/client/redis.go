package client

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/pkg/errors"
	cfg "qx/pkg/common/config"
)

// InitRedis ...
func InitRedis(config cfg.RedisConfiguration) (*redis.Client, error) {

	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return redisClient, nil
}

// CachedValue ...
func CachedValue(key string, notExistedCallback func() (string, error), redisCli *redis.Client, timeout time.Duration) (string, error) {
	var cachedValue string
	cachedValue, err := redisCli.Get(key).Result()
	if err == redis.Nil {
		cachedValue, err = notExistedCallback()
		if err != nil {
			return "", err
		}
		if timeout == 0 {
			timeout = 8 * time.Hour
		}
		err = redisCli.Set(key, cachedValue, timeout).Err()
		if err != nil {
			return "", errors.WithStack(err)
		}
	} else if err != nil {
		return "", errors.WithStack(err)
	}
	return cachedValue, nil
}
