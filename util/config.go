package util

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"GATEWAY_PORT"`
	AuthSVCUrl string `mapstructure:"AUTH_SVC_URL"`
	PostSVCUrl string `mapstructure:"POST_SVC_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
