package config

import (
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	config := viper.GetViper()
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AddConfigPath("/var/www/go")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Init Config File Err")
	}
}
func Get(key string) interface{} {
	return viper.Get(key)
}
func GetString(key string) string {
	return viper.GetString(key)
}
func GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
