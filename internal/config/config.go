package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	CompanyName string `mapstructure:"SEC_COMPANY_NAME"`
	CompanyEmail string `mapstructure:"SEC_COMPANY_EMAIL"`

}

func init() {
	// TODO: Add some config stuff
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Fprint(os.Stderr, err)
	// }
}


