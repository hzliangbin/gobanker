package options

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/hzliangbin/gobanker/pkg/apiserver"
	"github.com/hzliangbin/gobanker/pkg/blog"
	"github.com/hzliangbin/gobanker/pkg/database"
)

const (
	defaultConfiguration = "banker"
	defaultConfigPath = "/etc/banker"
)

type BankerOptions struct {
	DataBaseOpts *database.Config
	LoggerOpts *blog.Config
	APIServerOpts *apiserver.Config
}

func NewBankerOptions() *BankerOptions {
	bankerOpts := &BankerOptions{
		DataBaseOpts: &database.Config{},
		LoggerOpts: &blog.Config{},
		APIServerOpts: &apiserver.Config{},
	}
	return bankerOpts
}

func LoadConfigFromDisk() (*BankerOptions, error) {
	viper.SetConfigName(defaultConfiguration)
	viper.AddConfigPath(defaultConfigPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		} else {
			return nil, fmt.Errorf("error parsing configuration file %s", err)
		}
	}

	conf := NewBankerOptions()

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func (o *BankerOptions) Validate() []error {
	var errs []error

	return errs
}