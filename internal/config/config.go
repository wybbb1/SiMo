package config

import (
	"log"

	"github.com/spf13/viper"
)

var Config *AppConfig

type AppConfig struct {
	Server struct {
		Port string `yml:"Port"`
	} `yml:"Server"`

	Database struct {
		Type  string `yml:"Type"`
		Mysql struct {
			Host     string `yml:"Host"`
			Port     string `yml:"Port"`
			User     string `yml:"User"`
			Password string `yml:"Password"`
			DBName   string `yml:"DBName"`
		} `yml:"Mysql"`
		Sqlite struct {
			Path string `yml:"Path"`
		} `yml:"Sqlite"`
	} `yml:"Database"`

	Jwt struct {
		SecretKey     string `yml:"SecretKey"`
		ExpireHours   int    `yml:"ExpireHours"`
		Issuer        string `yml:"Issuer"`
		Subject       string `yml:"Subject"`
		Audience      string `yml:"Audience"`
	} `yml:"Jwt"`

	OAuth struct {
		ClientID     string `yml:"ClientID"`
		ClientSecret string `yml:"ClientSecret"`
		RedirectURL  string `yml:"RedirectURL"`
		Scopes       []string `yml:"Scopes"`
	} `yml:"OAuth"`

	Logging struct {
		Mode       string `yml:"Mode"`
		Path     []string `yml:"Path"`
		TimeLayout string `yml:"TimeLayout"`
	} `yml:"Logging"`
}

func InitConfig() { 
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	Config = &AppConfig{}
	err = viper.Unmarshal(Config)
	if err != nil {
		log.Fatalf("unable to decode into struct: %v", err)
	}

	log.Println("Configuration loaded successfully")
}
