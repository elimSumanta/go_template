package middleware

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"
)

type RedisCallwrapper struct {
	sf       *singleflight.Group
	rds      *redis.Client
	redisTTL int // in seconds
}

func InitRedisCallwrapper(rds *redis.Client, redisTTL int) *RedisCallwrapper {
	sf := &singleflight.Group{}
	return &RedisCallwrapper{sf, rds, redisTTL}
}

func (rds *RedisCallwrapper) CallWithRedisCache(ctx context.Context, strData interface{}, key string, fn func() (interface{}, error)) (interface{}, error) {
	res, err := rds.call(ctx, strData, key, func() (interface{}, error) {
		res, err, _ := rds.sf.Do(key, fn)
		return res, err
	})
	return res, err
}

func (rds *RedisCallwrapper) call(ctx context.Context, strData interface{}, key string, fn func() (interface{}, error)) (interface{}, error) {
	var (
		json = jsoniter.ConfigCompatibleWithStandardLibrary
	)
	if reflect.ValueOf(strData).Type().Kind() != reflect.Ptr {
		return nil, fmt.Errorf("data must be a pointer")
	}

	if rds.hasRedis() && strData != nil {
		val, err := rds.rds.Get(ctx, key).Result()
		if err == nil {
			errMarshall := json.UnmarshalFromString(val, strData)
			if errMarshall == nil {
				return strData, nil
			}
		}
	}

	res, err := fn()

	if err != nil || res == nil {
		return res, err
	}

	if rds.hasRedis() && key != "" {
		val, _ := json.MarshalToString(res)
		errRedis := rds.rds.SetEx(ctx, key, val, time.Duration(rds.redisTTL)*time.Second).Err()
		if errRedis != nil {
			logger.Errorf("[CallWrapper] Failed to set data on redis: %v", errRedis)
		}
	}

	return res, err
}

func (rds *RedisCallwrapper) hasRedis() bool {
	return rds.rds != nil && rds.redisTTL > 0
}
