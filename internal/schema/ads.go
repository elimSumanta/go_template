package schema

import "time"

// Variable Constants --------------------------------
type AdsType string
type PriceType string
type StatusAds string
type StatusTransaction string
type VerifyStatus string

// Constants --------------------------------
const (
	SELL AdsType = "sell"
	BUY  AdsType = "buy"

	FIXED    PriceType = "fixed"
	FLOATING PriceType = "floating"

	PUBLISHED   StatusAds = "published"
	UNPUBLISHED StatusAds = "unpublished"
	FINISHED    StatusAds = "finished"
	CANCELLED   StatusAds = "cancelled"

	INPROGRESS           StatusTransaction = "inprogress"
	CANCELLEDTRANSACTION StatusTransaction = "cancelled"
	SUCCESS              StatusTransaction = "success"

	AUTH  VerifyStatus = "auth"
	EMAIL VerifyStatus = "email"
	PHONE VerifyStatus = "phone"
	NONE  VerifyStatus = "none"
)

// Variable Structs ----------------------------------------------------------------
type AdsRepo struct {
	ID               int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	P2PUserID        string    `gorm:"column:p2p_users_id;type:varchar(255);NOT NULL"`
	Type             AdsType   `gorm:"column:'type';type:enum('sell','buy');NOT NULL"`
	Instrument       string    `gorm:"column:instrument;type:varchar(100);NOT NULL"`
	PriceType        PriceType `gorm:"column:'price_type';type:enum('fixed','floating');NOT NULL"`
	Price            int64     `gorm:"column:price;type:bigint(30);NOT NULL"`
	TotalAmount      int64     `gorm:"column:total_amount;type:bigint(30);NOT NULL"`
	OrderMax         int64     `gorm:"column:order_max;type:bigint(30);NOT NULL"`
	OrderMin         int64     `gorm:"column:order_min;type:bigint(30);NOT NULL"`
	PaymentMethodIds string    `gorm:"column:payment_method_ids;type:varchar(100);NOT NULL"` // list or payment method id splite by `,`
	Term             string    `gorm:"column:term;type:varchar(255);NULL"`
	AutoReplyMessage string    `gorm:"column:auto_replay_message;type:varchar(255);NULL"`
	RegionalID       int8      `gorm:"column:regional_id;type:int(11);NOT NULL"`
	ActivatedOTC     int       `gorm:"column:activated_otc;type:int(1);NOT NULL"`
	Status           StatusAds `gorm:"column:status;type:enum('published','unpublished','finished','cancelled');NOT NULL"`
	CreateTime       time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime       time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *AdsRepo) TableName() string {
	return "p2p_ads"
}

type AdsTransactionRepo struct {
	ID                 int64             `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	P2PAdsID           int64             `gorm:"column:p2p_ads_id;type:bigint(20);NOT NULL"`
	UserUUID           string            `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	Instrument         string            `gorm:"column:instrument;type:varchar(100);NOT NULL"`
	PriceFrom          int64             `gorm:"column:price_from;type:bigint(30);NOT NULL"`
	PriceTo            int64             `gorm:"column:price_to;type:bigint(30);NOT NULL"`
	StartedTime        time.Time         `gorm:"column:started_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	PaidTime           time.Time         `gorm:"column:paid_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	EndedTime          time.Time         `gorm:"column:ended_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime         time.Time         `gorm:"column:updated_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	Status             StatusTransaction `gorm:"column:'status';type:enum('inprogress','cancel','success');NOT NULL"`
	RegPaymentMethodId int64             `gorm:"column:reg_payment_method_id;type:bigint(20);NOT NULL"`
}

func (m *AdsTransactionRepo) TableName() string {
	return "p2p_transaction"
}

type AdsTransactionHistoryRepo struct {
	ID                    int64             `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	IDTransaction         int64             `gorm:"column:id_p2p_transaction;type:bigint(20);NOT NULL"`
	UserUUID              string            `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	Instrument            string            `gorm:"column:instrument;type:varchar(100);NOT NULL"`
	Price                 int64             `gorm:"column:price_from;type:bigint(30);NOT NULL"`
	Status                StatusTransaction `gorm:"column:'status';type:enum('create','inprogress','paid','cancel','success');NOT NULL"`
	VerificationStatus    VerifyStatus      `gorm:"column:'status';type:enum('auth','email','phone','none');NOT NULL"`
	EmailVerificationCode string            `gorm:"column:email_verification_code;type:varchar(255);NULL"`
	PhoneVerificationCode string            `gorm:"column:phone_verification_code;type:varchar(255);NULL"`
	CancelationStatusID   int64             `gorm:"column:cancelation_status_id;type:bigint(20);NULL"`
	CacelationReason      string            `gorm:"column:cancelation_reason;type:varchar(255);NULL"`
	CreatedTime           time.Time         `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *AdsTransactionHistoryRepo) TableName() string {
	return "p2p_transaction_history"
}

type CancelationStatusRepo struct {
	ID          int64     `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	Type        string    `gorm:"column:instrument;type:varchar(255);NOT NULL"`
	Reason      string    `gorm:"column:instrument;reason:varchar(255);NOT NULL"`
	CreatedTime time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *CancelationStatusRepo) TableName() string {
	return "p2p_cancelation_status"
}

// ----------------------------------------------------------------
