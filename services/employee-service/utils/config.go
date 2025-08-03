package utils

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Environment                string        `mapstructure:"ENVIRONMENT"`
	DBSource                   string        `mapstructure:"DB_SOURCE"`
	HTTPServerAddress          string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	SSEAddress                 string        `mapstructure:"SSEAddress"`
	MigrationUrl               string        `mapstructure:"MIGRATION_URL"`
	GRPCServerAddress          string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey          string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration        time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration       time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	DeviceAccessTokenDuration  time.Duration `mapstructure:"DEVICE_ACCESS_TOKEN_DURATION"`
	DeviceRefreshTokenDuration time.Duration `mapstructure:"DEVICE_REFRESH_TOKEN_DURATION"`
	DeepfaceServiceAddress     string        `mapstructure:"DEEPFACE_SERVICE_ADDRESS"`
}

func LoadConfig() (config Config, err error) {
	value, ok := os.LookupEnv("GESTY_ENVIRONMENT")
	if ok {
		return LoadCustomConfig(value)
	}

	return LoadCustomConfig("dev")
}

func LoadCustomConfig(environment string) (config Config, err error) {
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")
	viper.SetConfigName(environment)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
