package config

import (
	"path"
	"runtime"

	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")

	_, filename, _, _ := runtime.Caller(0)

	viper.AddConfigPath(path.Dir(filename))

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return
}
