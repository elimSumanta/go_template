package ads

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
)

func (r *AdsRepo) SetAds(ctx context.Context, req schema.AdsRepo) (int64, error) {
	result := r.db.WithContext(ctx).Create(&req)
	return result.RowsAffected, result.Error
}
