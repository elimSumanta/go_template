package feedback

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
)

func (r *FeedbackRepo) SetFeedBackByP2PUserID(ctx context.Context, req schema.FeedbackRepo) (int64, error) {
	result := r.db.WithContext(ctx).Create(&req)
	return result.RowsAffected, result.Error
}
