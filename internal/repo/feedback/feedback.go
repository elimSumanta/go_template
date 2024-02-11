package feedback

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type FeedbackRepo struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewFeedbackRepo(db *gorm.DB, rds *redis.Client) *FeedbackRepo {
	return &FeedbackRepo{db, rds}
}
