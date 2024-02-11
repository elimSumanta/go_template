package mainflow

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/model"
)

type MainFlowHandler struct {
	adsUsecase AdsUsecase
}

func Init(adsUsecase AdsUsecase) *MainFlowHandler {
	return &MainFlowHandler{adsUsecase}
}

type AdsUsecase interface {
	GetAdsList(ctx context.Context) ([]model.GetAdsListRes, error)
}
