package helper

import (
	"fmt"
	"log"

	library "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/library"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/schema"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	gormSchema "gorm.io/gorm/schema"
)

var datetimePrecision = 0

type DbInstance struct {
	*gorm.DB
}

func ConnMysql() (*DbInstance, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", library.GetConfig.DbUser, library.GetConfig.DbPassword, library.GetConfig.DbHost, library.GetConfig.DbPort, library.GetConfig.DbName)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultDatetimePrecision: &datetimePrecision,
	}), &gorm.Config{
		NamingStrategy: gormSchema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return &DbInstance{db}, nil
}

func (client *DbInstance) MySQLMigration() {
	err := client.AutoMigrate(
		&schema.AdsRepo{},
		&schema.AdsTransactionHistoryRepo{},
		&schema.AdsTransactionRepo{},
		&schema.CancelationStatusRepo{},
		&schema.FeedbackRepo{},
		&schema.P2PUsersRepo{},
		&schema.P2PUsersFollowersRepo{},
		&schema.P2PUsersOTCRepo{},
		&schema.PaymentMethodRepo{},
	)
	if err != nil {
		log.Fatalln("🔴 Database migration failed", err)
	}
}
