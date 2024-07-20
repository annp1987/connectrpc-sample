package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	ServicePort string
	DataFile    string
}

func NewConfig() *Config {
	confer := viper.New()
	confer.AutomaticEnv()
	confer.SetEnvPrefix("")
	confer.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// set default service port is 8080
	confer.SetDefault("port", "8080")
	confer.SetDefault("data.file", "sample.json")
	c := &Config{
		ServicePort: confer.GetString("port"),
		DataFile:    confer.GetString("data.file"),
	}
	return c
}
