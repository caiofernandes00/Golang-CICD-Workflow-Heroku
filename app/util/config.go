package util

import (
	"time"

	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port                 string        `mapstructure:"PORT"`
	MaxConcurrentStreams uint32        `mapstructure:"MAX_CONCURRENT_STREAMS"`
	MaxReadFrameSize     uint32        `mapstructure:"MAX_READ_FRAME_SIZE"`
	IdleTimeout          time.Duration `mapstructure:"IDLE_TIMEOUT"`
}

func LoadEnvFile(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	fmt.Printf("Error unmarshalling config, %s", err)
	return
}

func LoadEnv() (config Config, err error) {
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	fmt.Printf("Error unmarshalling config, %s", err)
	return
}
