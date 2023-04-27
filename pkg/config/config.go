package config

import (
	"github.com/spf13/viper"
)

type DBConfig struct {
	DbHost     string `mapstructure:"host"`
	DbName     string `mapstructure:"dbname"`
	DbPort     string `mapstructure:"port"`
	DbUser     string `mapstructure:"user"`
	DbPaswword string `mapstructure:"password"`
}

type JWTConfig struct {
	Secret_key string `mapstructure:"secret_key"`
}

type TwilioConfig struct {
	Service_id  string `mapstructure:"service_id"`
	Account_sid string `mapstructure:"accsid"`
	Auth_token  string `mapstructure:"authtoken"`
}

type Config struct {
	Db     DBConfig     `mapstructure:"db"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	Twilio TwilioConfig `mapstructure:"twilio"`
}

var config Config

func LoadConfig() (Config, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.SetConfigType("json")
	vp.AddConfigPath("/home/ashiq/Documents/shiftSync/pkg/config/")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func JwtConfig() string {
	return config.JWT.Secret_key
}
