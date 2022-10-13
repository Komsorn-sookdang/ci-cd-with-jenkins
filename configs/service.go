package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func Initialize() {
	viper.SetConfigName("configs")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// TODO: To be decided?
		} else {
			panic(fmt.Errorf("fatal error config file: %s", err))
		}
	}
}
