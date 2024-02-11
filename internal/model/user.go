package model

import "time"

// Variable Constants --------------------------------

// Constants --------------------------------

// Variable Structs ----------------------------------------------------------------
type UserProfile struct {
	UserID        string        `json:"user_id"`
	UserName      string        `json:"user_name"`
	UserStatus    UserStatus    `json:"user_status"`
	UserTradeInfo UserTradeInfo `json:"user_trade_info"`
	PaymentMethod string        `json:"payment_method"`
}

type UserStatus struct {
	IsVerified bool `json:"verified"`
	IsEmail    bool `json:"email"`
	IsSMS      bool `json:"sms"`
	ISKyc      bool `json:"is_kyc"`
}

type UserTradeInfo struct {
	TradeAvg           float32 `json:"trade_avg"`
	AllTimeTrade       int     `json:"all_time_trade"`
	CompletionsRateAvg float32 `json:"completions_rate_avg"`
	AllCompletionsRate int     `json:"all_completion_rate"`
	ReleaseTimeAvg     int     `json:"release_time_avg"`
	PayTimeAvg         float32 `json:"pay_time_avg"`
	LastTrade          int     `json:"last_trade"`
	ReceivedFeedBack   float32 `json:"received_feedback"`
}

type P2PUsersRepo struct {
	ID             int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	Name           string    `gorm:"column:name;type:varchar(255);NOT NULL"`
	UserUUID       string    `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	VerifiedEmail  bool      `gorm:"column:verified_email;type:tinyint(1);NOT NULL"`
	VerifiedMobile bool      `gorm:"column:verified_mobile;type:tinyint(1);NOT NULL"`
	VerifiedKYC    bool      `gorm:"column:verified_kyc;type:tinyint(1);NOT NULL"`
	CreatedTime    time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime     time.Time `gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *P2PUsersRepo) TableName() string {
	return "p2p_users"
}

type P2PUsersOTCRepo struct {
	ID           int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	P2PUserID    int64     `gorm:"column:id;type:bigint(20);NULL"`
	PosisionLong string    `gorm:"column:name;type:varchar(255);NOT NULL"`
	PosisionLat  string    `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	CreatedTime  time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime   time.Time `gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *P2PUsersOTCRepo) TableName() string {
	return "p2p_users_otc"
}

type P2PUsersFollowersRepo struct {
	ID          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	UserUUID    string    `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	P2PUserID   int64     `gorm:"column:id;type:bigint(20);NULL"`
	Notify      bool      `gorm:"column:notify;type:tinyint(1);NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime  time.Time `gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *P2PUsersFollowersRepo) TableName() string {
	return "p2p_followers"
}

// ----------------------------------------------------------------
