package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetInt64(key string) int64
	GetBool(key string) bool
	GetStringSlice(key string) []string
	Init()
}

type viperConfig struct {
	configFile string
}

func (v *viperConfig) Init() {
	replacer := strings.NewReplacer(`.`, `_`)
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType(`json`)
	viper.SetConfigFile(v.configFile)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

func (v *viperConfig) GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

func (v *viperConfig) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}

func NewViperConfig(configFile string) Config {
	v := &viperConfig{configFile}
	v.Init()
	return v
}
