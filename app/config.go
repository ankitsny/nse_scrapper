package app

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/spf13/viper"
)

type appConfig struct {
	ServerPort int    `mapstructure:"server_port"`
	LogLevel   string `mapstructure:"log_level"`
	LogPath    string `mapstructure:"log_path"`
}

// Config : Application configuration
var Config appConfig

// Validate : Validation function to validate the required params
func (c appConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ServerPort, validation.Required),
		validation.Field(&c.LogPath, validation.Required),
	)
}

// LoadConfig : Load config files
func LoadConfig(path string) error {
	v := viper.New()

	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		fmt.Println("ERR", err)
		return err
	}

	if err := v.Unmarshal(&Config); err != nil {
		return err
	}

	return Config.Validate()
}
