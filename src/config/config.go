package config

import (
	"github.com/spf13/viper"
)

type Config struct{}

func InitConfig() {
	// TODO: Add some config stuff
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Fprint(os.Stderr, err)
	// }
}

