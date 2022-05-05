package config

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"strings"
)

var (
	ContextKeyConfig = struct{}{}
)

type Config struct {
	AWS `mapstructure:"aws"`
}

type AWS struct {
	AccessKeyId string `mapstructure:"accessKey"`
	SecretKey   string `mapstructure:"secretKey"`
	Bucket      string `mapstructure:"bucket"`
}

func Load() (Config, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath("../../config")

	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := v.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func (c Config) SetConfig(ctx context.Context) context.Context {
	return context.WithValue(ctx, ContextKeyConfig, c)
}

func ReadConfig(ctx context.Context) (Config, error) {
	conf := ctx.Value(ContextKeyConfig)
	c, ok := conf.(Config)
	if !ok {
		return Config{}, errors.New("cannot get config struct")
	}
	return c, nil
}
