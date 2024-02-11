package schema

import "time"

// Variable Constants --------------------------------

// Constants --------------------------------

// Variable Structs ----------------------------------------------------------------
type PaymentMethodRepo struct {
	ID              uint      `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	UserUUID        string    `gorm:"column:user_uuid;type:varchar(255);NOT NULL"`
	PaymentMethodId int64     `gorm:"column:payment_method_id;type:bigint(20);NOT NULL"`
	Name            string    `gorm:"column:name;type:varchar(255);NOT NULL"`
	Number          string    `gorm:"column:number;type:varchar(255);NOT NULL"`
	Email           string    `gorm:"column:email;type:varchar(255);NOT NULL"`
	QRImage         string    `gorm:"column:qr_code_img;type:varchar(255);NOT NULL"`
	IsActive        bool      `gorm:"column:is_active;type:tinyint(1);NOT NULL"`
	CreateTime      time.Time `gorm:"column:created_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdateTime      time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (m *PaymentMethodRepo) TableName() string {
	return "p2p_reg_payment_method"
}

// ----------------------------------------------------------------
