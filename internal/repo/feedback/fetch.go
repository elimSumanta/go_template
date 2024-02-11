package feedback

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
)

func (r *FeedbackRepo) GetFeedBackByP2PUserID(ctx context.Context, p2pUserID int64) ([]schema.FeedbackRepo, error) {
	var data []schema.FeedbackRepo

	err := r.db.WithContext(ctx).Where("p2p_users_id = ?", p2pUserID).Find(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}
