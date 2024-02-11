package ads

import (
	"context"
	"fmt"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
	"gorm.io/gorm"
)

func (r *AdsRepo) GetAdsList(ctx context.Context, Status []schema.StatusAds) ([]schema.AdsRepo, error) {
	var dataStr []schema.AdsRepo

	key := fmt.Sprintf("ads:%+v", Status)
	res, err := r.callByRedis.GetAdsListCall.CallWithRedisCache(ctx, &dataStr, key, func() (interface{}, error) {
		var data []schema.AdsRepo
		err := r.db.WithContext(ctx).Where("status IN ?", Status).Find(&data).Error
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return data, err

	})

	if err != nil {
		return nil, err
	}

	data, ok := res.(*[]schema.AdsRepo)
	if ok {
		return dataStr, fmt.Errorf("failed to casting data to type %+v", res)
	}

	return *data, nil
}

func (r *AdsRepo) GetAdsListByMaker(ctx context.Context, p2pUserID int64, Status []string) ([]schema.AdsRepo, error) {
	var dataStr []schema.AdsRepo

	key := fmt.Sprintf("ads:%d:%+v", p2pUserID, Status)
	res, err := r.callByRedis.GetAdsListCall.CallWithRedisCache(ctx, &dataStr, key, func() (interface{}, error) {
		var data []schema.AdsRepo
		err := r.db.WithContext(ctx).Where("p2p_users_id = ? and status IN ?", p2pUserID, Status).Find(&data).Error
		if err == gorm.ErrRecordNotFound {
			err = nil
		}
		return data, err

	})

	if err != nil {
		return nil, err
	}

	data, ok := res.(*[]schema.AdsRepo)
	if ok {
		return dataStr, fmt.Errorf("failed to casting data to type %+v", res)
	}

	return *data, nil
}
