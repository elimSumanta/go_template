package schema

import "time"

// Variable Constants --------------------------------

// Constants --------------------------------

// Variable Structs ----------------------------------------------------------------
type FeedbackRepo struct {
	ID         uint      `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	P2PUserID  string    `gorm:"column:p2p_users_id;type:varchar(255);NOT NULL"`
	UserUUID   string    `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	Feedback   int8      `gorm:"column:feedback;type:int(8);NOT NULL"` // -1,0,1
	Message    string    `gorm:"column:message;type:varchar(255);NOT NULL"`
	CreateTime time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *FeedbackRepo) TableName() string {
	return "p2p_feedback"
}

// ----------------------------------------------------------------
