package config

import (
	"log"

	"github.com/spf13/viper"
)

var c *viper.Viper

func Init(env string) {
	c = viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName(env)
	c.AddConfigPath("config/environment")
	if err := c.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return c
}
