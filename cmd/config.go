package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct{}

func initConfig() {
	// TODO: Add some config stuff
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
