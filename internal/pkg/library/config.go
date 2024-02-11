package library

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ApplicationEnv struct {
	Env         string `mapstructure:"ENV"`
	EnableCors  bool   `mapstructure:"ENABLE_CORS"`
	BasicApiKey string `mapstructure:"BASIC_API_KEY"`
	SecretKey   string `mapstructure:"APP_SECRET"`

	HttpReadTimeOut  string `mapstructure:"HTTP_READ_TIMEOUT"`
	HttpWriteTimeOut string `mapstructure:"HTTP_WRITE_TIMEOUT"`

	DbHost      string `mapstructure:"SQL_HOST"`
	DbPort      string `mapstructure:"SQL_PORT"`
	DbUser      string `mapstructure:"SQL_USER"`
	DbPassword  string `mapstructure:"SQL_PASSWORD"`
	DbName      string `mapstructure:"SQL_DB"`
	DbMigration bool   `mapstructure:"SQL_AUTO_MIGRATION"`

	RdbHost string `mapstructure:"REDIS_HOST"`
	RdbPass string `mapstructure:"REDIS_PASS"`
}

var GetConfig = &ApplicationEnv{}

func InitAppConfig() {
	var configStruct = GetConfig
	viper.SetConfigFile(".env")
	readConfig(configStruct)
}

func readConfig(config *ApplicationEnv) {
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Error("Cannot read configuration", err)
	}

	setDefaultEnv()

	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Error("Cannot read configuration: ", err, ". will use default env")
	}
}

func setDefaultEnv() map[string]string {
	m := make(map[string]string)
	for _, s := range os.Environ() {
		a := strings.Split(s, "=")
		if viper.Get("TEMPLATE") == "true" {
			viper.Set(a[0], a[1])
		}
		m[a[0]] = a[1]
	}
	return m
}
