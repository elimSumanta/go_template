package cmd

import (
	"context"

	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/helper"
	library "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/library"
	"github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/logger"
	"github.com/joho/godotenv"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	_ "github.com/rookie-ninja/rk-gin/v2/boot"
	rkquery "github.com/rookie-ninja/rk-query"

	log "github.com/bitwyre/bitwyre/gateway/rest/p2p_api/internal/pkg/logger"
	"github.com/redis/go-redis/v9"
)

func Server() {
	_ = godotenv.Load()

	// Create a new boot instance.
	boot := rkboot.NewBoot()

	log.LoggerInit()

	// Init dependencies
	library.InitAppConfig()

	eventEntry := rkentry.GlobalAppCtx.GetEventEntry("my-event")
	event := eventEntry.CreateEvent(rkquery.WithOperation("test"))
	event.AddPair("key", "value")
	event.Finish()

	boot.Bootstrap(context.Background())
	InitApp()
	boot.WaitForShutdownSig(context.Background())

}

func InitApp() {
	// Connect to the database
	db, errDB := helper.ConnMysql()
	if errDB != nil {
		panic(errDB)
	}

	if library.GetConfig.DbMigration {
		db.MySQLMigration()
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     library.GetConfig.RdbHost,
		Password: library.GetConfig.RdbPass, // no password set
		DB:       0,                         // use default DB
	})

	// Init services
	services, err := NewService(db.DB, rdb)
	if err != nil {
		logger.Fatal(err.Error())
	}

	router := InitRouter(services)
	router.InitPrivate()
	router.InitPublic()

}
