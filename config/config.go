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
	JwtSecret         string
)

func init() {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	viper.SetDefault("PORT", 3040)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	Port = viper.GetInt("PORT")
	LocalDatabaseName = viper.GetString("LOCAL_DATABASE_NAME")
	DatabaseUrl = viper.GetString("DATABASE_URL")
	DatabaseAuthToken = viper.GetString("DATABASE_AUTH_TOKEN")
	JwtSecret = viper.GetString("JWT_SECRET")
}
