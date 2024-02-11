package ads

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
)

type AdsUsecase struct {
	repo RepoInitiated
}

type RepoInitiated struct {
	AdsRepo AdsRepo
}

func InitAdsUsecase(repo RepoInitiated) *AdsUsecase {
	return &AdsUsecase{repo}
}

type AdsRepo interface {
	GetAdsList(ctx context.Context, Status []schema.StatusAds) ([]schema.AdsRepo, error)
}
