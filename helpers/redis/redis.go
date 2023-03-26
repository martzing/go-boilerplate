package redisHelper

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/martzing/go-boilerplate/configs"
	utilityHelper "github.com/martzing/go-boilerplate/helpers/utility"
	"go.uber.org/zap"
)

type redisClient struct {
	client *redis.Client
	logger *zap.SugaredLogger
}

var redisConnection *redisClient

func New() *redisClient {
	redisConfig := configs.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", redisConfig.Host, redisConfig.Port),
		DB:   1,
	})

	logger := utilityHelper.NewLogger()

	if err := client.Ping(context.Background()).Err(); err != nil {
		logger.Fatalln(err)
	}

	logger.Info("Connected Redis")

	redisConnection = &redisClient{
		client: client,
		logger: logger,
	}

	return redisConnection
}

func GetRedisInstance() *redisClient {
	if redisConnection == nil {
		return New()
	}
	return redisConnection
}

func (r *redisClient) GetClient() *redis.Client {
	return r.client
}

func (r *redisClient) Set(key string, fields map[string]interface{}) bool {
	hash, err := r.client.HMSet(context.Background(), key, fields).Result()
	if err != nil {
		r.logger.Errorln(err)
	}
	return hash
}

func (r *redisClient) Get(key string, field string) ([]interface{}, error) {
	return r.client.HMGet(context.Background(), key, field).Result()
}

func (r *redisClient) SetExpire(key string, duration time.Duration) {
	r.client.Expire(context.Background(), key, duration)
}
