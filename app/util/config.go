package util

import (
	"reflect"
	"time"

	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Port                             string        `mapstructure:"PORT"`
	MaxConcurrentStreams             uint32        `mapstructure:"MAX_CONCURRENT_STREAMS"`
	MaxReadFrameSize                 uint32        `mapstructure:"MAX_READ_FRAME_SIZE"`
	IdleTimeout                      time.Duration `mapstructure:"IDLE_TIMEOUT"`
	CircuitBreakerInterval           time.Duration `mapstructure:"CIRCUIT_BREAKER_INTERVAL"`
	CircuitBreakerThreshold          int           `mapstructure:"CIRCUIT_BREAKER_THRESHOLD"`
	CacheRequestCapacity             int           `mapstructure:"CACHE_REQUEST_CAPACITY"`
	CacheRequestTTL                  time.Duration `mapstructure:"CACHE_REQUEST_TTL"`
	SkipCompressionUrls              []string      `mapstructure:"SKIP_COMPRESSION_URLS"`
	RateLimiter                      time.Duration `mapstructure:"RATE_LIMITER"`
	RateLimiterExponentialBaseFactor time.Duration `mapstructure:"RATE_LIMITER_EXPONENTIAL_BASE_FACTOR"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) LoadEnvFile(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&c)
	c.unmarshalDurationFields()
	fmt.Printf("Error unmarshalling config, %s", err)

	return
}

func (c *Config) LoadEnv() (config Config, err error) {
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	fmt.Printf("Error unmarshalling config, %s", err)
	return
}

func (c *Config) unmarshalDurationFields() {
	t := reflect.TypeOf(*c)

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i).Type

		if fieldType == reflect.TypeOf(time.Duration(0)) {
			value := viper.GetDuration(t.Field(i).Tag.Get("mapstructure"))
			reflect.ValueOf(c).Elem().Field(i).Set(reflect.ValueOf(value * time.Second))
		}
	}
}
