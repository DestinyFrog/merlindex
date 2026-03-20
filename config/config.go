package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	Port              int
	LocalDatabaseName string
	DatabaseUrl       string
	DatabaseAuthToken string
)

func init() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 3040)
	viper.SetDefault("DATABASE_FILE", "purple.db")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	Port = viper.GetInt("PORT")
}
