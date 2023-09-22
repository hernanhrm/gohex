package config

import (
	"log"

	"github.com/spf13/viper"
)

type LocalConfig struct {
	ServerPort     uint   `mapstructure:"SERVER_PORT"`
	AllowedDomains string `mapstructure:"ALLOWED_DOMAINS"`
	AllowedMethods string `mapstructure:"ALLOWED_METHODS"`
	Database       Database
}

type Database struct {
	Driver   string `mapstructure:"DATABASE_DRIVER"`
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     uint   `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
}

func LoadLocalConfig() LocalConfig {
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("viper.ReadInConfig(): fatal error config file: %v", err)
	}

	var localConfig LocalConfig
	if err := viper.Unmarshal(&localConfig); err != nil {
		log.Fatalf("viper.Unmarshal(): %v", err)
	}

	if err := viper.Unmarshal(&localConfig.Database); err != nil {
		log.Fatalf("viper.Unmarshal(): %v", err)
	}

	return localConfig
}
