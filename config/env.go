package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"sync"
)

var (
	envOnce     sync.Once
	envInstance *viper.Viper
)

func Env() *viper.Viper {
	envOnce.Do(func() {
		envInstance = newEnv()
	})

	return envInstance
}

func newEnv() *viper.Viper {
	v := viper.New()
	v.SetConfigName("env")
	v.SetConfigType("yml")
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Read in config error: %v", err)
	}

	v.AutomaticEnv()

	return v
}
