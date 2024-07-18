package viper

import (
	"github.com/spf13/viper"
)

const CONFIG_YAML = "yaml"

func Viper(path, name string) error {
	viper.SetConfigName(name)
	viper.SetConfigType(CONFIG_YAML)
	viper.AddConfigPath(path)

	return viper.ReadInConfig()

}
