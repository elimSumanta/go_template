package ads

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/logger"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
)

func (uc *AdsUsecase) GetAdsList(ctx context.Context) (res []model.GetAdsListRes, err error) {
	_, err = uc.repo.AdsRepo.GetAdsList(ctx, []schema.StatusAds{schema.PUBLISHED})
	if err != nil {
		logger.Errorf("Error getting Ads list: %v", err)
	}

	return res, nil
}
