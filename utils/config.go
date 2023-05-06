package utils

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	TokenSecretKey        string        `mapstructure:"TOKEN_SECRET_KEY"`
	HTTPServerPort        string        `mapstructure:"HTTP_SERVER_PORT"`
	TokenDuration         time.Duration `mapstructure:"TOKEN_DURATION"`
	RefreshTokenSecretKey string        `mapstructure:"REFRESH_TOKEN_SECRET_KEY"`
	TokenRefreshDuration  time.Duration `mapstructure:"TOKEN_REFRESH_DURATION"`
	DSN                   string        `mapstructure:"DSN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
	})

	err = viper.Unmarshal(&config)
	return
}
