package ads

import (
	md "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/middleware"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AdsRepo struct {
	db          *gorm.DB
	rds         *redis.Client
	callByRedis CallByRedis
}

type CallByRedis struct {
	GetAdsListCall *md.RedisCallwrapper
}

func NewAdsRepo(db *gorm.DB, rds *redis.Client) *AdsRepo {
	callByRedis := CallByRedis{
		GetAdsListCall: md.InitRedisCallwrapper(rds, 2),
	}
	return &AdsRepo{db, rds, callByRedis}
}
